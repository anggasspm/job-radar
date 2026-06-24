import os
import glob
from dotenv import load_dotenv
from db_manager import get_db_connection

# Memuat variabel dari file .env (kalau ada) ke environment
load_dotenv()


def run_migrations():
    print("=== Memulai Inisialisasi Database ===")

    # Path menuju folder migrasi milik Backend
    migrations_dir = "../backend/migrations"

    conn = get_db_connection()
    if not conn:
        print("Migrasi dibatalkan karena koneksi database gagal.")
        return

    cursor = conn.cursor()

    try:
        # Mencari semua file berakhiran .up.sql dan mengurutkannya (000001, 000002, dst)
        sql_files = sorted(glob.glob(os.path.join(migrations_dir, "*.up.sql")))

        if not sql_files:
            print("Tidak menemukan file migrasi. Pastikan struktur foldernya benar.")
            return

        for file_path in sql_files:
            filename = os.path.basename(file_path)
            print(f"Mengeksekusi: {filename}...")

            with open(file_path, "r", encoding="utf-8") as f:
                sql_query = f.read()
                cursor.execute(sql_query)

        conn.commit()
        print("-> Sukses! Semua tabel dasar berhasil dibuat.")

    except Exception as e:
        print(f"Gagal menjalankan migrasi: {e}")
        conn.rollback()

    finally:
        cursor.close()
        conn.close()


if __name__ == "__main__":
    run_migrations()