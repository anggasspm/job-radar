def normalize_salary(salary_text):
    """
    TODO: Ekstrak min dan max salary dari string menggunakan regex.
    Kembalikan salary_min, salary_max sebagai integer.
    """
    return 10000000, 15000000

def normalize_job_data(raw_job):
    """
    Membersihkan 1 row data lowongan.
    """
    salary_min, salary_max = normalize_salary(raw_job.get("salary_text", ""))
    
    clean_job = {
        "title": raw_job.get("title"),
        "company": raw_job.get("company"),
        "location": raw_job.get("location").replace("Jkt", "Jakarta"), # Contoh normalisasi lokasi
        "salary_min": salary_min,
        "salary_max": salary_max,
        "category": "Software Engineering", # TODO: Buat logic pemetaan kategori
        "raw_url": raw_job.get("raw_url"),
        "source": raw_job.get("source")
    }
    return clean_job