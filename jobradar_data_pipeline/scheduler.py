import schedule
import time
import datetime
from main import run_pipeline

def job():
    """Fungsi pembungkus untuk menjalankan pipeline utama"""
    print(f"\n[{datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')}] Menjalankan JobRadar Pipeline Terjadwal...")
    try:
        run_pipeline()
        print(f"[{datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')}] Pipeline selesai dieksekusi.")
    except Exception as e:
        print(f"[{datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')}] ERROR pada pipeline: {e}")

def start_scheduler():
    print("=== JobRadar Scheduler Aktif ===")
    print("Menunggu waktu jadwal eksekusi...")
    
    # ---------------------------------------------------------
    # KONFIGURASI JADWAL
    # ---------------------------------------------------------
    # Opsi 1: Jalankan setiap hari jam 02:00 pagi (Standar Industri)
    # schedule.every().day.at("02:00").do(job)
    
    # Opsi 2: Jalankan setiap 12 jam sekali
    # schedule.every(12).hours.do(job)

    # Opsi 3: UNTUK TESTING (Jalankan setiap 1 menit)
    schedule.every(1).minutes.do(job)
    # ---------------------------------------------------------

    # (Opsional) Langsung jalankan 1x saat script pertama kali dihidupkan
    # job()

    # Loop abadi untuk terus mengecek waktu
    while True:
        schedule.run_pending()
        time.sleep(60) # Cek setiap 60 detik agar CPU tidak kerja keras

if __name__ == "__main__":
    start_scheduler()