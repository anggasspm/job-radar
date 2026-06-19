CREATE TABLE users (
    id              BIGSERIAL PRIMARY KEY,
    email           VARCHAR(255) NOT NULL UNIQUE,
    password_hash   VARCHAR(255),                  -- NULL kalau user cuma pernah login via OAuth (lihat oauth_accounts)
    name            VARCHAR(255),
    avatar_url      VARCHAR(500),
    tier            VARCHAR(10) NOT NULL DEFAULT 'free' CHECK (tier IN ('free', 'premium')),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);