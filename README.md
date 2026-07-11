# Job Radar

Job Radar adalah platform agregator lowongan kerja yang mengumpulkan data dari beberapa sumber (scraping), menormalisasi dan menyimpannya ke database, lalu menyajikannya melalui REST API dan antarmuka web modern. Pengguna bisa mencari, memfilter, dan menyimpan lowongan favorit.

## Fitur Utama

- **Pencarian & Filter Lowongan** — cari berdasarkan kata kunci, kategori, lokasi, dan rentang gaji.
- **Autentikasi Pengguna** — registrasi & login dengan JWT (access + refresh token).
- **Favorit** — simpan lowongan favorit per pengguna.
- **Data Pipeline (ETL)** — scraper otomatis dari beberapa sumber, dinormalisasi, lalu disimpan ke PostgreSQL, dijadwalkan lewat GitHub Actions.
- **Dokumentasi API** — Swagger/OpenAPI otomatis.
- **Mock API** — server mock (json-server style) untuk pengembangan frontend tanpa backend asli.

## Struktur Project

Repo ini adalah monorepo dengan 4 komponen utama:

```
job-radar/
├── backend/                 # REST API (Go + Gin + GORM + PostgreSQL)
├── jobradar-frontend/       # Web app (Next.js + React + Tailwind CSS)
├── jobradar-mock-api/       # Mock API server (Node.js/Express) untuk dev frontend
├── jobradar_data_pipeline/  # ETL pipeline scraper (Python)
└── .github/workflows/       # CI/CD & scheduler (scraper.yml)
```

## Tech Stack

| Layer | Teknologi |
|---|---|
| Backend API | Go, Gin, GORM, PostgreSQL (Neon), JWT, Swagger |
| Frontend | Next.js 16, React 19, Tailwind CSS 4, lucide-react |
| Data Pipeline | Python, scraping (GraphQL/HTML), scheduler |
| Mock API | Node.js |
| Infra | Docker, Docker Compose, GitHub Actions |

---

## Getting Started

### 1. Backend (Go API)

**Prasyarat:** Go 1.26+, Docker, PostgreSQL (atau gunakan docker-compose).

```bash
cd backend

# Jalankan database lokal via Docker
docker compose up -d

# Buat file .env berisi konfigurasi berikut:
# APP_PORT=8080
# DSN=host=localhost user=root password=root dbname=job-radar port=5450 sslmode=disable
# APP_SECRET=<secret-key-anda>

# Jalankan server (hot reload pakai Air)
make server
```

- Sinkronisasi schema database ke `database/schema.sql`:
  ```bash
  make schema
  ```
- Dokumentasi API (Swagger UI) tersedia di:
  ```
  http://localhost:8080/swagger/index.html
  ```

**Endpoint utama:**

| Method | Endpoint | Deskripsi | Auth |
|---|---|---|---|
| POST | `/auth/register` | Registrasi user baru | ❌ |
| POST | `/auth/login` | Login user | ❌ |
| GET | `/jobs/` | List semua lowongan | ❌ |
| GET | `/jobs/search` | Cari/filter lowongan | ❌ |
| GET | `/jobs/:id` | Detail lowongan | ❌ |
| GET | `/favorite/` | List lowongan favorit user | ✅ |

### 2. Frontend (Next.js)

**Prasyarat:** Node.js 18+.

```bash
cd jobradar-frontend
npm install
npm run dev
```

Aplikasi berjalan di `http://localhost:3000`. Pastikan variabel environment/API base URL diarahkan ke backend Go (`http://localhost:8080`) atau ke mock API saat development.

### 3. Mock API (opsional, untuk dev frontend tanpa backend Go)

```bash
cd jobradar-mock-api
npm install
node server.js
```

### 4. Data Pipeline (Python ETL)

**Prasyarat:** Python 3.10+.

```bash
cd jobradar_data_pipeline
pip install -r requirements.txt

# Buat file .env berisi kredensial database PostgreSQL

# Jalankan pipeline sekali (extract -> transform -> load)
python main.py

# Atau jalankan terjadwal
python scheduler.py
```

Pipeline ini melakukan scraping dari beberapa sumber (`scrapers/source_a.py`, `source_b.py`, `source_c.py`), menormalisasi data (`processor/normalizer.py`), lalu menyimpannya ke PostgreSQL (`database/db_manager.py`). Proses ini juga dijalankan otomatis terjadwal via GitHub Actions (`.github/workflows/scraper.yml`).

---

## Skema Database

Database menggunakan PostgreSQL dengan tabel utama:

- `users`, `oauth_accounts`, `refresh_tokens` — autentikasi & akun pengguna
- `jobs`, `sources`, `skills`, `job_skills` — data lowongan kerja
- `salary_snapshots` — statistik gaji historis
- `favorites` — lowongan favorit pengguna
- `api_keys`, `api_usage_daily` — manajemen API key & rate limit
- `alert_subscriptions`, `alert_notifications` — notifikasi lowongan baru

Migrasi database dikelola di `backend/database/migrations/` menggunakan format `golang-migrate`.

## Docker

Backend dapat dijalankan penuh via Docker:

```bash
cd backend
docker build -t job-radar-backend .
docker run -p 7860:7860 --env-file .env job-radar-backend
```

## Kontribusi

1. Fork repository ini.
2. Buat branch baru: `git checkout -b fitur/nama-fitur`.
3. Commit perubahan Anda.
4. Push dan buat Pull Request.

## Lisensi

Belum ditentukan.
