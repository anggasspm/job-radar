CREATE TABLE api_keys (
    id                  BIGSERIAL PRIMARY KEY,
    user_id             BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    key_hash            VARCHAR(255) NOT NULL UNIQUE,  -- simpan hash, JANGAN plaintext
    label               VARCHAR(100),
    rate_limit_per_day  INT NOT NULL DEFAULT 100,
    is_active           BOOLEAN NOT NULL DEFAULT true,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    last_used_at        TIMESTAMPTZ
);