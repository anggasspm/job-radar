def normalize_salary(amount):
    if amount is None:
        return 0
    try:
        return int(float(amount))
    except (ValueError, TypeError):
        return 0

def normalize_location(raw_location):
    if not raw_location:
        return "Tidak diketahui"
    loc = raw_location.title()
    prefixes_to_remove = ["Kota Administrasi ", "Kota ", "Kabupaten ", "Kab. "]
    for prefix in prefixes_to_remove:
        loc = loc.replace(prefix, "")
    return loc.strip()

def categorize_job(title):
    """
    Mengelompokkan judul pekerjaan mentah ke kategori baku berbasis kata kunci.
    """
    title_lower = title.lower()
    
    # Aturan pencocokan (dari yang spesifik ke umum)
    if any(kw in title_lower for kw in ["backend", "back end", "back-end", "golang", "java developer", "python developer", "node.js"]):
        return "Backend Developer"
    elif any(kw in title_lower for kw in ["frontend", "front end", "front-end", "react", "vue", "angular"]):
        return "Frontend Developer"
    elif any(kw in title_lower for kw in ["fullstack", "full stack", "full-stack"]):
        return "Fullstack Developer"
    elif any(kw in title_lower for kw in ["mobile", "ios", "android", "flutter", "react native", "kotlin"]):
        return "Mobile Developer"
    elif any(kw in title_lower for kw in ["data scientist", "data engineer", "data analyst", "machine learning", "ai "]):
        return "Data Professional"
    elif any(kw in title_lower for kw in ["devops", "sre", "site reliability", "cloud", "infrastructure"]):
        return "DevOps & Cloud"
    elif any(kw in title_lower for kw in ["ui/ux", "ui / ux", "product designer", "user experience"]):
        return "UI/UX Designer"
    elif any(kw in title_lower for kw in ["product manager", "scrum master", "business analyst", "project manager"]):
        return "Product & Project Management"
    elif any(kw in title_lower for kw in ["security", "pentester", "cyber", "grc"]):
        return "Cyber Security"
    
    return "Lainnya" # Kategori fallback jika tidak ada keyword yang cocok

def normalize_job_data(raw_job):
    raw_title = raw_job.get("title", "").strip()
    
    clean_job = {
        "title": raw_title,
        "company": raw_job.get("company", "").strip(),
        "location": normalize_location(raw_job.get("location")),
        "category": categorize_job(raw_title), # FUNGSI BARU DIPANGGIL DI SINI
        "min_salary": normalize_salary(raw_job.get("min_salary")),
        "max_salary": normalize_salary(raw_job.get("max_salary")),
        "currency": raw_job.get("currency") or "IDR",
        "min_exp": raw_job.get("min_exp") or 0,
        "max_exp": raw_job.get("max_exp") or 0,
        "education": raw_job.get("education", "").replace("_", " ").title(),
        "skills": raw_job.get("skills", []),
        "source": raw_job.get("source"),
        "raw_url": raw_job.get("raw_url")
    }
    
    return clean_job