CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE jobs (
    id              BIGSERIAL PRIMARY KEY,
    source_id       INT NOT NULL REFERENCES sources(id),
    external_id     VARCHAR(255),                -- id asli dari API source (job.id Glints/TIA), null utk WWR
    title           VARCHAR(255) NOT NULL,
    company         VARCHAR(255) NOT NULL,
    location        VARCHAR(255) NOT NULL,
    category        VARCHAR(100),                -- role yang sudah dinormalisasi, contoh: "Backend Developer".
                                                   -- BELUM ada di normalizer.py saat ini — lihat catatan di bawah.
    description     TEXT,                         -- belum di-scrape saat ini, disiapkan utk job detail page
    salary_min      BIGINT NOT NULL DEFAULT 0,
    salary_max      BIGINT NOT NULL DEFAULT 0,
    currency        VARCHAR(3) NOT NULL DEFAULT 'IDR',
    min_exp         SMALLINT NOT NULL DEFAULT 0,
    max_exp         SMALLINT NOT NULL DEFAULT 0,
    education       VARCHAR(50),
    raw_url         VARCHAR(500) NOT NULL UNIQUE, -- dedup key (sesuai TODO "ON CONFLICT" di db_manager.py)
    posted_date     DATE,
    scraped_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    is_active       BOOLEAN NOT NULL DEFAULT true, -- false kalau lowongan sudah hilang dari source saat re-scrape
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_jobs_category_location ON jobs(category, location);
CREATE INDEX idx_jobs_salary             ON jobs(salary_min, salary_max);
CREATE INDEX idx_jobs_scraped_at         ON jobs(scraped_at);
CREATE INDEX idx_jobs_source             ON jobs(source_id);
CREATE INDEX idx_jobs_active             ON jobs(is_active) WHERE is_active = true;
-- full text search sederhana utk fallback non-AI search / debugging
CREATE INDEX idx_jobs_title_trgm ON jobs USING gin (title gin_trgm_ops);