CREATE TABLE experiment_users (
    id BIGSERIAL PRIMARY KEY,
    experiment_id BIGINT,
    group_id BIGINT,
    user_id BIGINT,
    anonymous_id VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (experiment_id) REFERENCES experiments(id),
    FOREIGN KEY (group_id) REFERENCES experiment_groups(id)
);

CREATE INDEX exp_users_idx_user_id ON experiment_users (user_id);
CREATE INDEX exp_users_idx_anonymous_id ON experiment_users (anonymous_id);
