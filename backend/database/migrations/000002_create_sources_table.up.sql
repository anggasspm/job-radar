CREATE TABLE sources (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL UNIQUE,     -- "Glints", "Tech in Asia", "We Work Remotely"
    base_url    VARCHAR(255),
    is_active   BOOLEAN NOT NULL DEFAULT true,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);


INSERT INTO sources (name, base_url) VALUES
    ('Glints', 'https://glints.com'),
    ('Tech in Asia', 'https://www.techinasia.com'),
    ('We Work Remotely', 'https://weworkremotely.com');
 