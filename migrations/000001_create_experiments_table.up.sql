CREATE TABLE experiments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT FALSE,
    stores JSON,
    platforms JSON,
    only_authorized BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX exp_idx_is_active ON experiments (is_active);
CREATE INDEX exp_idx_slug ON experiments (slug);
CREATE INDEX exp_idx_only_authorized ON experiments (only_authorized);
