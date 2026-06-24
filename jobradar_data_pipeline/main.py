from dotenv import load_dotenv
load_dotenv()
from scrapers.source_a import scrape_glints_graphql
from scrapers.source_b import scrape_source_b
from scrapers.source_c import scrape_source_c
from processor.normalizer import normalize_job_data
from database.db_manager import save_jobs  # Import fungsi save_jobs

def run_pipeline():
    print("=== Memulai JobRadar Data Pipeline ===")
    
    # --- TAHAP 1: EXTRACT ---
    print("\n[1/3] Memulai proses ekstraksi (Extract)...")
    
    raw_data_glints = scrape_glints_graphql()
    raw_data_tia = scrape_source_b()
    raw_data_wwr = scrape_source_c()
    
    # Menggabungkan semua data mentah menjadi satu list
    all_raw_data = raw_data_glints + raw_data_tia + raw_data_wwr
    
    if not all_raw_data:
        print("Tidak ada data yang berhasil diekstrak. Pipeline dihentikan.")
        return
        
    print(f"-> Selesai: Total {len(all_raw_data)} lowongan mentah ditarik.")
    
    # --- TAHAP 2: TRANSFORM ---
    print("\n[2/3] Memulai proses normalisasi (Transform)...")
    clean_data = []
    
    for raw_job in all_raw_data:
        cleaned = normalize_job_data(raw_job)
        clean_data.append(cleaned)
        
    print(f"-> Selesai: {len(clean_data)} lowongan berhasil dinormalisasi.")
    
    # --- TAHAP 3: LOAD ---
    print("\n[3/3] Memulai penyimpanan ke Database PostgreSQL (Load)...")
    
    # Menyimpan data bersih langsung ke database PostgreSQL
    save_jobs(clean_data)
    
    print("\n=== Pipeline Selesai dengan Sukses ===")

if __name__ == "__main__":
    run_pipeline()