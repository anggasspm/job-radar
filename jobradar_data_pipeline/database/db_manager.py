import psycopg2

def get_db_connection():
    # Kredensial ini disesuaikan dengan isi docker-compose.yml backend
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
    """Menyimpan data ke tabel jobs dengan deduplikasi."""
    if not clean_jobs_list:
        return

    print(f"\nMemulai proses Load {len(clean_jobs_list)} lowongan ke PostgreSQL...")
    conn = get_db_connection()
    if not conn:
        return

    cursor = conn.cursor()
    inserted_count = 0

    # Mapping nama source ke source_id sesuai migrasi backend (000002_create_sources_table.up.sql)
    source_map = {
        "Glints": 1,
        "Tech in Asia": 2,
        "We Work Remotely": 3
    }

    # Query INSERT dengan klausa ON CONFLICT DO NOTHING untuk mencegah duplikasi (berdasarkan raw_url)
    insert_query = """
        INSERT INTO jobs (
            source_id, title, company, location, salary_min, salary_max, 
            currency, min_exp, max_exp, education, raw_url
        ) VALUES (
            %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s
        ) ON CONFLICT (raw_url) DO NOTHING;
    """

    for job in clean_jobs_list:
        source_id = source_map.get(job['source'], 1) # Default ke 1 jika tidak ditemukan
        
        try:
            cursor.execute(insert_query, (
                source_id,
                job['title'],
                job['company'],
                job['location'],
                job['min_salary'],
                job['max_salary'],
                job['currency'],
                job['min_exp'],
                job['max_exp'],
                job['education'],
                job['raw_url']
            ))
            # Jika rowcount == 1, berarti data baru berhasil masuk (bukan duplikat)
            if cursor.rowcount == 1:
                inserted_count += 1
                
        except Exception as e:
            print(f"Error menyimpan job {job['title']}: {e}")
            conn.rollback() # Rollback jika ada error pada satu row agar loop bisa lanjut

    conn.commit()
    cursor.close()
    conn.close()
    
    print(f"-> Selesai: {inserted_count} lowongan baru berhasil ditambahkan. (Sisanya adalah duplikat yang diabaikan).")