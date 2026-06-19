CREATE TABLE alert_subscriptions (
    id              BIGSERIAL PRIMARY KEY,
    user_id         BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    keyword         VARCHAR(255),
    category        VARCHAR(100),
    location        VARCHAR(255),
    min_salary      BIGINT,
    channel         VARCHAR(10) NOT NULL CHECK (channel IN ('email', 'telegram')),
    channel_target  VARCHAR(255) NOT NULL,   -- alamat email atau chat_id Telegram
    is_active       BOOLEAN NOT NULL DEFAULT true,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_alerts_active ON alert_subscriptions(is_active) WHERE is_active = true;