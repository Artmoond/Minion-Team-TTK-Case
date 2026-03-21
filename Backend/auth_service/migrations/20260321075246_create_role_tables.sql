-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL PRIMARY KEY,
    rolename TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE IF EXISTS roles;
