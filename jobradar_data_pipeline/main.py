from scrapers.source_a import scrape_source_a
from processor.normalizer import normalize_job_data
from database.db_manager import save_jobs

def run_pipeline():
    print("--- Memulai JobRadar Data Pipeline ---")
    
    # 1. Ekstrak
    raw_data_a = scrape_source_a()
    # raw_data_b = scrape_source_b()
    
    all_raw_data = raw_data_a # + raw_data_b
    
    # 2. Transformasi / Normalisasi
    clean_data = []
    for data in all_raw_data:
        cleaned = normalize_job_data(data)
        clean_data.append(cleaned)
        
    # 3. Load / Simpan
    if clean_data:
        save_jobs(clean_data)
        
    print("--- Pipeline Selesai ---")

if __name__ == "__main__":
    # Script ini bisa dieksekusi oleh Cron Job setiap beberapa jam
    run_pipeline()