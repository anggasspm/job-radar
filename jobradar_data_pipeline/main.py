from scrapers.source_a import scrape_source_a
from processor.normalizer import normalize_job_data
from database.db_manager import save_jobs

def run_pipeline():
    print("--- Memulai Uji Coba JobRadar Data Pipeline ---")
    
    # 1. Ekstrak data dari Sumber A
    raw_data = scrape_source_a()
    
    # Tampilkan hasil sementara
    print(f"Data yang berhasil ditarik: {len(raw_data)} item")
    if raw_data:
        print(raw_data[0]) # Cek isi data pertama

if __name__ == "__main__":
    run_pipeline()

if __name__ == "__main__":
    # Script ini bisa dieksekusi oleh Cron Job setiap beberapa jam
    run_pipeline()

