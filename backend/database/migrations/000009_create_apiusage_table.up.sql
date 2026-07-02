CREATE TABLE api_usage_daily (
    api_key_id      BIGINT NOT NULL REFERENCES api_keys(id) ON DELETE CASCADE,
    usage_date      DATE NOT NULL,
    request_count   INT NOT NULL DEFAULT 0,
    PRIMARY KEY (api_key_id, usage_date)
);