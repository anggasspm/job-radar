import os
import psycopg2

def get_db_connection():
    try:
        # Pendekatan 1: Coba baca format URL tunggal (Sering dipakai platform Cloud/Go)
        # Contoh: postgresql://user:pass@host:port/dbname
        db_url = os.getenv("DATABASE_URL")
        if db_url:
            conn = psycopg2.connect(db_url)
            return conn
        
        # Pendekatan 2: Coba baca 5 variabel terpisah (Sesuai dengan GitHub Secrets sebelumnya)
        db_host = os.getenv("DB_HOST")
        db_port = os.getenv("DB_PORT")
        db_name = os.getenv("DB_NAME")
        db_user = os.getenv("DB_USER")
        db_pass = os.getenv("DB_PASSWORD")

        # Pastikan kelima variabel ini memiliki nilai (tidak None)
        if all([db_host, db_port, db_name, db_user, db_pass]):
            conn = psycopg2.connect(
                host=db_host,
                port=db_port,
                dbname=db_name,
                user=db_user,
                password=db_pass
            )
            return conn
        
        # Jika kedua pendekatan di atas gagal
        print("Gagal: Environment variable database belum diatur dengan benar.")
        print("Pastikan DATABASE_URL atau (DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASSWORD) sudah terisi.")
        return None

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

    # QUERY 1: Insert Job
    job_insert_query = """
        INSERT INTO jobs (
            source_id, title, company, location, category, description, salary_min, salary_max, 
            currency, min_exp, max_exp, education, raw_url
        ) VALUES (
            %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s
        ) ON CONFLICT (raw_url) DO UPDATE 
          SET description = EXCLUDED.description,
              updated_at = now()
        RETURNING id;
    """

    # QUERY 2: Insert Skill
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
                job['category'], job['description'],
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

            # 2. Proses Relasi Skills
            unique_skills = set(job.get('skills', []))

            for skill_name in unique_skills:
                clean_skill = skill_name.strip()
                if not clean_skill:
                    continue

                # Coba masukkan skill baru
                cursor.execute(skill_insert_query, (clean_skill,))
                skill_record = cursor.fetchone()

                if skill_record:
                    skill_id = skill_record[0]
                    new_skills_added += 1
                else:
                    cursor.execute("SELECT id FROM skills WHERE name = %s", (clean_skill,))
                    skill_id = cursor.fetchone()[0]

                # 3. Hubungkan Job ID dengan Skill ID
                cursor.execute(job_skill_insert_query, (job_id, skill_id))

            # Commit per lowongan agar jika 1 lowongan error, tidak membatalkan semuanya
            conn.commit()

        except Exception as e:
            print(f"Error menyimpan job '{job['title']}': {e}")
            conn.rollback()

    cursor.close()
    conn.close()

    print(f"-> Selesai: {inserted_jobs} data pekerjaan telah diproses (Insert/Update).")
    print(f"-> Info: Terdapat {new_skills_added} entry skill baru ke dalam database kamus skill.")