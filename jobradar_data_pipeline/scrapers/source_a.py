import requests
import xml.etree.ElementTree as ET
import gzip
import io

def scrape_glints_sitemap():
    print("Mulai membaca Sitemap Index Glints...")
    
    # URL Sitemap Index
    index_url = "https://glints.com/sitemaps/explore-page-id-sitemap.xml" 
    
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
    }
    
    try:
        # --- LANGKAH 1: Ambil Sitemap Index ---
        response = requests.get(index_url, headers=headers)
        response.raise_for_status()
        
        root_index = ET.fromstring(response.content)
        namespace = {'ns': 'http://www.sitemaps.org/schemas/sitemap/0.9'}
        
        # Ambil URL sitemap pertama dari index (untuk testing MVP)
        first_sitemap = root_index.find('ns:sitemap', namespace)
        if first_sitemap is None:
            print("Tidak menemukan tag sitemap di dalam index.")
            return []
            
        gz_url = first_sitemap.find('ns:loc', namespace).text
        print(f"Menemukan target sitemap terkompresi:\n-> {gz_url}")
        
        # --- LANGKAH 2: Unduh dan Ekstrak file .gz ---
        print("\nMengunduh dan mengekstrak file .gz...")
        gz_response = requests.get(gz_url, headers=headers)
        gz_response.raise_for_status()
        
        # Membaca data byte terkompresi dan mengekstraknya
        with gzip.GzipFile(fileobj=io.BytesIO(gz_response.content)) as f:
            xml_content = f.read()
            
        # --- LANGKAH 3: Parsing XML yang sudah diekstrak ---
        root = ET.fromstring(xml_content)
        
        job_urls = []
        limit = 5
        count = 0
        
        for url_tag in root.findall('ns:url', namespace):
            if count >= limit:
                break
                
            loc = url_tag.find('ns:loc', namespace).text
            job_urls.append(loc)
            count += 1
                
        print(f"\nBerhasil mengekstrak {len(job_urls)} URL dari sub-sitemap.")
        return job_urls
        
    except requests.exceptions.RequestException as e:
        print(f"Error HTTP: {e}")
    except ET.ParseError as e:
        print(f"Error parsing XML: {e}")
    except Exception as e:
        print(f"Terjadi error yang tidak terduga: {e}")
        
    return []

if __name__ == "__main__":
    urls = scrape_glints_sitemap()
    print("\n--- Hasil Ekstraksi URL ---")
    for u in urls:
        print(u)