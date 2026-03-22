-- +goose Up
CREATE TABLE IF NOT EXISTS user_roles (
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    role_id BIGINT REFERENCES roles(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS user_roles;
