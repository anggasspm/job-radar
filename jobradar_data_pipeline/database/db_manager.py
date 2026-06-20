import psycopg2

def get_db_connection():
    try:
        conn = psycopg2.connect(
            host="localhost",
            port="5450",
            dbname="job-radar",
            user="root",
            password="root"
        )
        return conn
    except Exception as e:
        print(f"Gagal terhubung ke database: {e}")
        return None

def save_jobs(clean_jobs_list):
    """Menyimpan data ke tabel jobs beserta relasi skill-nya."""
    if not clean_jobs_list:
        return

    print(f"\nMemulai proses Load {len(clean_jobs_list)} lowongan ke PostgreSQL (beserta relasi Skills)...")
    conn = get_db_connection()
    if not conn:
        return

    cursor = conn.cursor()
    inserted_jobs = 0
    new_skills_added = 0

    source_map = {
        "Glints": 1,
        "Tech in Asia": 2,
        "We Work Remotely": 3
    }

    # QUERY 1: Insert Job. 
    # Kita gunakan trik DO UPDATE SET updated_at = now() agar Postgres SELALU mengembalikan ID
    job_insert_query = """
        INSERT INTO jobs (
            source_id, title, company, location, salary_min, salary_max, 
            currency, min_exp, max_exp, education, raw_url
        ) VALUES (
            %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s
        ) ON CONFLICT (raw_url) DO UPDATE 
          SET updated_at = now()
        RETURNING id;
    """

    # QUERY 2: Insert Skill (Hanya jika belum ada)
    skill_insert_query = """
        INSERT INTO skills (name) VALUES (%s) 
        ON CONFLICT (name) DO NOTHING 
        RETURNING id;
    """

    # QUERY 3: Insert Relasi Job & Skill
    job_skill_insert_query = """
        INSERT INTO job_skills (job_id, skill_id) VALUES (%s, %s)
        ON CONFLICT (job_id, skill_id) DO NOTHING;
    """

    for job in clean_jobs_list:
        source_id = source_map.get(job['source'], 1)
        
        try:
            # 1. Eksekusi penyimpanan Job dan tangkap ID-nya
            cursor.execute(job_insert_query, (
                source_id, job['title'], job['company'], job['location'],
                job['min_salary'], job['max_salary'], job['currency'],
                job['min_exp'], job['max_exp'], job['education'], job['raw_url']
            ))
            
            job_record = cursor.fetchone()
            if not job_record:
                # Fallback jaga-jaga jika DO UPDATE tidak jalan
                cursor.execute("SELECT id FROM jobs WHERE raw_url = %s", (job['raw_url'],))
                job_record = cursor.fetchone()
                
            job_id = job_record[0]
            inserted_jobs += 1

            # 2. Proses Relasi Skills (gunakan set() untuk menghilangkan skill duplikat dalam 1 job)
            unique_skills = set(job.get('skills', []))
            
            for skill_name in unique_skills:
                clean_skill = skill_name.strip()
                if not clean_skill:
                    continue
                    
                # Coba masukkan skill baru
                cursor.execute(skill_insert_query, (clean_skill,))
                skill_record = cursor.fetchone()
                
                if skill_record:
                    # Skill baru berhasil ditambahkan
                    skill_id = skill_record[0]
                    new_skills_added += 1
                else:
                    # Skill sudah ada di database, kita cukup SELECT ID-nya
                    cursor.execute("SELECT id FROM skills WHERE name = %s", (clean_skill,))
                    skill_id = cursor.fetchone()[0]
                    
                # 3. Hubungkan Job ID dengan Skill ID
                cursor.execute(job_skill_insert_query, (job_id, skill_id))
            
            # Commit per lowongan agar jika 1 lowongan error, tidak membatalkan semuanya
            conn.commit()
            
        except Exception as e:
            print(f"Error menyimpan job '{job['title']}': {e}")
            conn.rollback() # Batalkan transaksi spesifik ini

    cursor.close()
    conn.close()
    
    print(f"-> Selesai: {inserted_jobs} data pekerjaan telah diproses (Insert/Update).")
    print(f"-> Info: Terdapat {new_skills_added} entry skill baru ke dalam database kamus skill.")