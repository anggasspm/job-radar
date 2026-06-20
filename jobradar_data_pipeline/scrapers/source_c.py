import requests
import xml.etree.ElementTree as ET

def scrape_source_c():
    print("Mulai mengambil data dari RSS Feed We Work Remotely...")
    
    # URL RSS Feed WWR khusus kategori Backend/Programming
    url = "https://weworkremotely.com/categories/remote-back-end-programming-jobs.rss"
    
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
    }
    
    raw_jobs = []
    
    try:
        response = requests.get(url, headers=headers)
        response.raise_for_status()
        
        # Membaca format XML dari RSS Feed
        root = ET.fromstring(response.content)
        
        # Dalam struktur RSS, data ada di dalam <channel> lalu di dalam <item>
        items = root.findall('./channel/item')
        print(f"Berhasil menarik {len(items)} lowongan dari We Work Remotely!\n")
        
        for item in items:
            # Format judul di RSS WWR selalu "Nama Perusahaan: Judul Pekerjaan"
            full_title = item.findtext('title') or ""
            
            # Perbaikan Logika Split untuk menghindari 'https:'
            if ":" in full_title:
                parts = full_title.split(":", 1)
                # Mengecek apakah kata sebelum ':' adalah 'http' atau 'https'
                if parts[0].strip().lower() in ["http", "https"]:
                    company_name = "Perusahaan tidak diketahui"
                    job_title = full_title
                else:
                    company_name = parts[0]
                    job_title = parts[1]
            else:
                company_name = "Perusahaan tidak diketahui"
                job_title = full_title
                
            # WWR berfokus pada remote, jadi lokasi default adalah Remote
            location = "Remote / Anywhere"
            link = item.findtext('link') or ""
            
            # --- MASUKKAN KE FORMAT STANDAR NORMALIZER ---
            raw_jobs.append({
                "title": job_title.strip(),
                "company": company_name.strip(),
                "location": location,
                "description": item.findtext('description'),
                "min_salary": 0,    # RSS WWR jarang mencantumkan gaji di metadata
                "max_salary": 0,
                "currency": "USD",  # Mayoritas menggunakan USD
                "min_exp": 0,
                "max_exp": 0,
                "education": "Tidak disebutkan",
                "skills": [],       # Data skill tercampur di deskripsi HTML
                "source": "We Work Remotely",
                "raw_url": link.strip(),
                "raw_data": {}      # Kita kosongkan karena datanya cukup rata (flat)
            })
            
            print(f"- {job_title.strip()} | {company_name.strip()} ({location})")
            
    except requests.exceptions.RequestException as e:
        print(f"Error HTTP saat mengakses RSS WWR: {e}")
    except ET.ParseError as e:
        print(f"Error parsing XML RSS: {e}")
        
    return raw_jobs

if __name__ == "__main__":
    scrape_source_c()