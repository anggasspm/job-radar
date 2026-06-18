import requests
from bs4 import BeautifulSoup

def scrape_source_a():
    """
    Fungsi ini khusus mengambil data mentah dari Sumber A.
    Mengembalikan list of dictionaries berisi raw data.
    """
    raw_jobs = []
    # TODO: Implementasi requests.get() dan BeautifulSoup disini
    # Contoh struktur dummy:
    # raw_jobs.append({
    #     "title": "Backend Engineer",
    #     "company": "PT Inovasi",
    #     "location": "Jkt",
    #     "salary_text": "Rp 10.000.000 - 15.000.000",
    #     "raw_url": "https://...",
    #     "source": "Source_A"
    # })
    print("Mengeksekusi scraper Sumber A...")
    return raw_jobs