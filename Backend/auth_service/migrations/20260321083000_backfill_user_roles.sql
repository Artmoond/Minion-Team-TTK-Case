-- +goose Up
INSERT INTO roles (rolename)
VALUES ('Пользователь')
ON CONFLICT (rolename) DO NOTHING;

INSERT INTO roles (rolename)
SELECT DISTINCT role
FROM users
WHERE role IS NOT NULL
ON CONFLICT (rolename) DO NOTHING;

INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id
FROM users u
JOIN roles r ON r.rolename = u.role
WHERE u.role IS NOT NULL
  AND NOT EXISTS (
      SELECT 1
      FROM user_roles ur
      WHERE ur.user_id = u.id
        AND ur.role_id = r.id
  );

DELETE FROM user_roles ur
USING user_roles dup
WHERE ur.ctid < dup.ctid
  AND ur.user_id = dup.user_id
  AND ur.role_id = dup.role_id;

ALTER TABLE user_roles
    ALTER COLUMN user_id TYPE BIGINT,
    ALTER COLUMN role_id TYPE BIGINT,
    ALTER COLUMN user_id SET NOT NULL,
    ALTER COLUMN role_id SET NOT NULL;

CREATE UNIQUE INDEX IF NOT EXISTS user_roles_user_id_role_id_idx
    ON user_roles (user_id, role_id);

ALTER TABLE users
    DROP COLUMN IF EXISTS role;

-- +goose Down
ALTER TABLE users
    ADD COLUMN IF NOT EXISTS role TEXT;

UPDATE users u
SET role = COALESCE(
    (
        SELECT r.rolename
        FROM user_roles ur
        JOIN roles r ON r.id = ur.role_id
        WHERE ur.user_id = u.id
        ORDER BY r.id
        LIMIT 1
    ),
    'Пользователь'
);

DROP INDEX IF EXISTS user_roles_user_id_role_id_idx;
