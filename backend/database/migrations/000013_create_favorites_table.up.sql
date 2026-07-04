CREATE TABLE favorites (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    job_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_favorites_user
        FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_favorites_job
        FOREIGN KEY (job_id) REFERENCES jobs(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_favorite UNIQUE (user_id, job_id)
);