def normalize_salary(amount):
    """Memastikan gaji menjadi integer atau 0 jika kosong/error."""
    if amount is None:
        return 0
    try:
        return int(float(amount))
    except (ValueError, TypeError):
        return 0

def normalize_location(raw_location):
    """Membersihkan format lokasi kota/kabupaten."""
    if not raw_location:
        return "Tidak diketahui"
    
    # Ubah ke Title Case (contoh: "KOTA JAKARTA" -> "Kota Jakarta")
    loc = raw_location.title()
    
    # Hapus prefix administratif agar lebih seragam untuk pencarian
    prefixes_to_remove = ["Kota Administrasi ", "Kota ", "Kabupaten ", "Kab. "]
    for prefix in prefixes_to_remove:
        loc = loc.replace(prefix, "")
        
    return loc.strip()

def normalize_job_data(raw_job):
    """
    Fungsi utama untuk menormalisasi 1 baris data lowongan kerja.
    Menerima dictionary raw_job dari scraper.
    """
    clean_job = {
        "title": raw_job.get("title", "").strip(),
        "company": raw_job.get("company", "").strip(),
        "location": normalize_location(raw_job.get("location")),
        "min_salary": normalize_salary(raw_job.get("min_salary")),
        "max_salary": normalize_salary(raw_job.get("max_salary")),
        "currency": raw_job.get("currency") or "IDR",
        "min_exp": raw_job.get("min_exp") or 0,
        "max_exp": raw_job.get("max_exp") or 0,
        # Mengubah "BACHELOR_DEGREE" menjadi "Bachelor Degree"
        "education": raw_job.get("education", "").replace("_", " ").title(),
        "skills": raw_job.get("skills", []),
        "source": raw_job.get("source"),
        "raw_url": raw_job.get("raw_url")
    }
    
    return clean_job