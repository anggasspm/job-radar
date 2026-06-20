import requests
import json

def scrape_source_b():
    print("Mulai mengambil data dari API Algolia (Tech in Asia)...")
    
    # URL Endpoint Algolia. Parameter ID dan Key diletakkan langsung di query URL.
    url = "https://219wx3mpv4-dsn.algolia.net/1/indexes/*/queries?x-algolia-agent=Algolia%20for%20vanilla%20JavaScript%203.30.0%3BJS%20Helper%202.26.1&x-algolia-application-id=219WX3MPV4&x-algolia-api-key=b528008a75dc1c4402bfe0d8db8b3f8e"
    
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
        "Content-Type": "application/x-www-form-urlencoded",
        "Accept": "application/json"
    }
    
    # Format request Algolia. Perhatikan facetFilters untuk memfilter Indonesia
    payload = {
        "requests": [
            {
                "indexName": "job_postings",
                "params": "query=&hitsPerPage=30&page=0&facetFilters=[[\"city.country_name:Indonesia\"]]"
            }
        ]
    }
    
    raw_jobs = []
    
    try:
        response = requests.post(url, headers=headers, json=payload)
        response.raise_for_status()
        
        data = response.json()
        
        # Algolia mengembalikan data di dalam list 'results', di dalamnya ada list 'hits'
        jobs_list = data.get("results", [])[0].get("hits", [])
        
        print(f"Berhasil menarik {len(jobs_list)} lowongan dari Tech in Asia!\n")
        
        for job in jobs_list:
            # Ekstraksi data dari JSON Algolia
            title = job.get("title") or "Tidak ada judul"
            company = job.get("company", {}).get("name") or "Perusahaan tidak diketahui"
            
            # Lokasi
            city_name = job.get("city", {}).get("name")
            country_name = job.get("city", {}).get("country_name")
            location = f"{city_name}, {country_name}" if city_name else (country_name or "Lokasi tidak diketahui")
            
            # Gaji & Mata Uang
            min_salary = job.get("salary_min")
            max_salary = job.get("salary_max")
            currency = job.get("currency", {}).get("currency_code")
            
            # Pengalaman
            min_exp = job.get("experience_min")
            max_exp = job.get("experience_max")
            
            # Skills (Bentuk aslinya adalah list of dictionaries dengan key 'name')
            skills_data = job.get("job_skills") or []
            skill_names = [s.get("name") for s in skills_data if s.get("name")]
            
            # --- MASUKKAN KE FORMAT STANDAR ---
            raw_jobs.append({
                "title": title,
                "company": company,
                "location": location,
                "description": job.get("description"),
                "min_salary": min_salary,
                "max_salary": max_salary,
                "currency": currency,
                "min_exp": min_exp,
                "max_exp": max_exp,
                "education": "Tidak disebutkan", # TIA jarang menyebutkan pendidikan di level teratas
                "skills": skill_names,
                "source": "Tech in Asia",
                "raw_url": f"https://www.techinasia.com/jobs/{job.get('id')}",
                "raw_data": job
            })
            
            print(f"- {title} | {company} ({location})")
            
    except requests.exceptions.RequestException as e:
        print(f"Error HTTP saat mengakses API Tech in Asia: {e}")
    except json.JSONDecodeError:
        print("Gagal membaca struktur JSON dari response Tech in Asia.")
    except Exception as e:
        print(f"Error tidak terduga pada scraper TIA: {e}")
        
    return raw_jobs

if __name__ == "__main__":
    results = scrape_source_b()
    print("\nContoh struktur raw_job:")
    if results:
        print(json.dumps(results[0], indent=2))