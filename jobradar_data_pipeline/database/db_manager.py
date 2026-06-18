import psycopg2 # Pastikan library psycopg2 terinstall

def get_db_connection():
    # TODO: Implementasi koneksi ke PostgreSQL
    pass

def save_jobs(clean_jobs_list):
    """
    Menyimpan data ke tabel jobs. 
    TODO: Gunakan ON CONFLICT (raw_url) DO NOTHING untuk deduplikasi.
    """
    print(f"Menyimpan {len(clean_jobs_list)} lowongan ke database...")
    # conn = get_db_connection()
    # cursor = conn.cursor()
    # for job in clean_jobs_list:
    #     cursor.execute("INSERT INTO jobs (...) VALUES (...) ON CONFLICT ...")
    # conn.commit()