CREATE TABLE salary_snapshots (
    id              BIGSERIAL PRIMARY KEY,
    category        VARCHAR(100) NOT NULL,
    location        VARCHAR(255) NOT NULL,
    avg_salary      NUMERIC(14,2) NOT NULL,
    min_salary      BIGINT,
    max_salary      BIGINT,
    job_count       INT NOT NULL,
    snapshot_date   DATE NOT NULL,
    UNIQUE (category, location, snapshot_date)
);
 
CREATE INDEX idx_snapshots_lookup ON salary_snapshots(category, location, snapshot_date);