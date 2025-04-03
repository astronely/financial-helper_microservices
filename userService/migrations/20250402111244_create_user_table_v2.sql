-- +goose Up
ALTER TABLE users
    ALTER COLUMN email TYPE VARCHAR(256),
    ADD CONSTRAINT users_email_unique UNIQUE (email),
    ALTER COLUMN name TYPE VARCHAR(128),
    ALTER COLUMN password TYPE VARCHAR(256),
    ADD COLUMN password_changed_at TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    ADD COLUMN photo VARCHAR(512),
    ALTER COLUMN created_at TYPE TIMESTAMPTZ,
    ALTER COLUMN updated_at SET DEFAULT NOW();

-- +goose Down
ALTER TABLE users
    DROP COLUMN password_changed_at,
    DROP COLUMN photo,
    ALTER COLUMN email TYPE TEXT,
    DROP CONSTRAINT users_email_unique,
    ALTER COLUMN name TYPE TEXT,
    ALTER COLUMN password TYPE TEXT,
    ALTER COLUMN created_at TYPE TIMESTAMP;