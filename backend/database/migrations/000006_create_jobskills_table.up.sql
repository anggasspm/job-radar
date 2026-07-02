CREATE TABLE job_skills (
    job_id      BIGINT NOT NULL REFERENCES jobs(id) ON DELETE CASCADE,
    skill_id    INT NOT NULL REFERENCES skills(id) ON DELETE CASCADE,
    PRIMARY KEY (job_id, skill_id)
);

CREATE INDEX idx_job_skills_skill ON job_skills(skill_id);