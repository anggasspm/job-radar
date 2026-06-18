from scrapers.source_a import scrape_glints_graphql
from processor.normalizer import normalize_job_data
import json

def run_pipeline():
    print("=== Memulai JobRadar Data Pipeline ===")
    
    # --- TAHAP 1: EXTRACT ---
    print("\n[1/3] Memulai proses ekstraksi (Extract)...")
    raw_data_glints = scrape_glints_graphql()
    
    # Jika menggunakan fungsi scraper lama yang nge-print banyak hal, 
    # output terminalnya mungkin agak ramai. Tidak apa-apa untuk MVP.
    
    if not raw_data_glints:
        print("Tidak ada data yang berhasil diekstrak. Pipeline dihentikan.")
        return
        
    print(f"-> Selesai: {len(raw_data_glints)} lowongan mentah ditarik.")
    
    # --- TAHAP 2: TRANSFORM ---
    print("\n[2/3] Memulai proses normalisasi (Transform)...")
    clean_data = []
    for raw_job in raw_data_glints:
        cleaned = normalize_job_data(raw_job)
        clean_data.append(cleaned)
        
    print(f"-> Selesai: {len(clean_data)} lowongan berhasil dinormalisasi.")
    
    # Menampilkan 1 contoh hasil data bersih
    print("\n=== CONTOH 1 DATA BERSIH ===")
    print(json.dumps(clean_data[0], indent=2))
    
    # --- TAHAP 3: LOAD (Segera Hadir) ---
    print("\n[3/3] Proses penyimpanan ke Database (Load) - Menunggu setup database.")

if __name__ == "__main__":
    run_pipeline()