CREATE TABLE experiment_groups (
    id SERIAL PRIMARY KEY,
    experiment_id BIGINT,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    is_active BOOLEAN DEFAULT FALSE,
    coverage_percentage FLOAT CHECK (coverage_percentage >= 0 AND coverage_percentage <= 100),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (experiment_id) REFERENCES experiments(id)
);

CREATE INDEX exp_groups_idx_is_active ON experiment_groups (is_active);
