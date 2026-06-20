import requests
import json

def scrape_glints_graphql():
    print("Mulai mengambil data dari GraphQL API Glints...")
    
    # URL persis seperti yang kamu temukan di tab Network
    url = "https://glints.com/api/v2-alc/graphql?op=searchJobsV3"
    
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
        "Content-Type": "application/json",
        "Accept": "application/json"
    }
    
    # Payload yang sudah disesuaikan dengan sintaks dictionary Python
    payload = {
        "operationName": "searchJobsV3",
        "variables": {
            "data": {
                "SearchTerm": "Backend Developer",
                "CountryCode": "ID",
                "pageSize": 30,
                "page": 1, # Saya ubah ke 1 agar kita mengambil halaman pertama
                "includeExternalJobs": True # Python menggunakan huruf kapital untuk boolean
            }
        },
        # PERUBAHAN 1: Menambahkan 'description\n' di dalam query GraphQL
        "query": "query searchJobsV3($data: JobSearchConditionInput!) {\n  searchJobsV3(data: $data) {\n    jobsInPage {\n      id\n      title\n      description\n      workArrangementOption\n      status\n      createdAt\n      updatedAt\n      isHot\n      isApplied\n      shouldShowSalary\n      educationLevel\n      type\n      fraudReportFlag\n      company {\n        ...CompanyFields\n        __typename\n      }\n      citySubDivision {\n        id\n        name\n        __typename\n      }\n      city {\n        ...CityFields\n        __typename\n      }\n      country {\n        ...CountryFields\n        __typename\n      }\n      salaries {\n        ...SalaryFields\n        __typename\n      }\n      location {\n        ...LocationFields\n        __typename\n      }\n      minYearsOfExperience\n      maxYearsOfExperience\n      source\n      jobSource\n      type\n      hierarchicalJobCategory {\n        id\n        level\n        name\n        children {\n          name\n          level\n          id\n          __typename\n        }\n        parents {\n          id\n          level\n          name\n          __typename\n        }\n        __typename\n      }\n      skills {\n        skill {\n          id\n          name\n          __typename\n        }\n        mustHave\n        __typename\n      }\n      traceInfo\n      __typename\n    }\n    expInfo\n    hasMore\n    __typename\n  }\n}\n\nfragment CompanyFields on Company {\n  id\n  name\n  brandName\n  logo\n  status\n  isVIP\n  IndustryId\n  industry {\n    id\n    name\n    __typename\n  }\n  verificationTier {\n    type\n    userName\n    __typename\n  }\n  __typename\n}\n\nfragment CityFields on City {\n  id\n  name\n  __typename\n}\n\nfragment CountryFields on Country {\n  code\n  name\n  __typename\n}\n\nfragment SalaryFields on JobSalary {\n  id\n  salaryType\n  salaryMode\n  maxAmount\n  minAmount\n  CurrencyCode\n  __typename\n}\n\nfragment LocationFields on HierarchicalLocation {\n  id\n  name\n  administrativeLevelName\n  formattedName\n  level\n  slug\n  latitude\n  longitude\n  parents {\n    id\n    name\n    administrativeLevelName\n    formattedName\n    level\n    slug\n    CountryCode: countryCode\n    parents {\n      level\n      formattedName\n      slug\n      __typename\n    }\n    __typename\n  }\n  __typename\n}"
    }
    
    raw_jobs = []
    
    try:
        response = requests.post(url, headers=headers, json=payload)
        response.raise_for_status()
        
        data = response.json()
        
        # Menavigasi struktur JSON hasil kembalian GraphQL
        # Urutan: data -> searchJobsV3 -> jobsInPage (berupa list of dictionaries)
        jobs_list = data.get("data", {}).get("searchJobsV3", {}).get("jobsInPage", [])
        
        print(f"Berhasil menarik {len(jobs_list)} lowongan!\n")
        
        for job in jobs_list:
            title = job.get("title") or "Tidak ada judul"
            
            company_data = job.get("company") or {}
            company = company_data.get("name") or "Perusahaan tidak diketahui"
            
            # --- WATERFALL LOCATION EXTRACTION ---
            location_name = None
            
            # Prioritas 1: Ambil dari object 'location'
            loc_data = job.get("location") or {}
            if loc_data.get("name"):
                location_name = loc_data.get("name")
                
            # Prioritas 2: Jika kosong, coba dari 'city'
            if not location_name:
                city_data = job.get("city") or {}
                location_name = city_data.get("name")
                
            # Prioritas 3: Jika masih kosong, coba dari 'country'
            if not location_name:
                country_data = job.get("country") or {}
                location_name = country_data.get("name")
                
            # Fallback jika semua object di atas null
            location = location_name or "Lokasi tidak diketahui"

            # --- EKSTRAKSI GAJI ---
            salaries_list = job.get("salaries") or []
            if salaries_list:
                salary_info = salaries_list[0]
                min_salary = salary_info.get("minAmount")
                max_salary = salary_info.get("maxAmount")
                currency = salary_info.get("CurrencyCode")
            else:
                min_salary = None
                max_salary = None
                currency = None

            # --- EKSTRAKSI REQUIREMENTS ---
            min_exp = job.get("minYearsOfExperience")
            max_exp = job.get("maxYearsOfExperience")
            edu_level = job.get("educationLevel") or "Tidak disebutkan"
            
            skills_data = job.get("skills") or []
            skill_names = []
            for s in skills_data:
                skill_info = s.get("skill") or {}
                skill_name = skill_info.get("name")
                if skill_name:
                    skill_names.append(skill_name)
            
            # --- UPDATE DICTIONARY ---
            raw_jobs.append({
                "title": title,
                "company": company,
                "location": location,
                "description": job.get("description"), # PERUBAHAN 2: Menangkap data description
                "min_salary": min_salary,
                "max_salary": max_salary,
                "currency": currency,
                "min_exp": min_exp,
                "max_exp": max_exp,
                "education": edu_level,
                "skills": skill_names,
                "source": "Glints",
                "raw_url": f"https://glints.com/id/opportunities/jobs/{job.get('id')}",
                "raw_data": job
            })
            
            skills_str = ", ".join(skill_names) if skill_names else "Tidak ada spesifikasi skill"
            print(f"- {title} | {company}")
            print(f"  Lokasi: {location} | Exp: {min_exp}-{max_exp} thn | Edu: {edu_level}")
            print(f"  Skills: {skills_str[:60]}...\n")

    except requests.exceptions.RequestException as e:
        print(f"Error HTTP saat mengakses API: {e}")
    except json.JSONDecodeError:
        print("Gagal membaca struktur JSON dari response.")
        
    return raw_jobs

if __name__ == "__main__":
    scrape_glints_graphql()