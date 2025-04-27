-- +goose Up
-- +goose StatementBegin
CREATE TYPE "roles" AS ENUM (
  'admin',
  'editor',
  'viewer'
);

CREATE TABLE "boards"
(
    "id"          bigserial PRIMARY KEY,
    "name"        varchar(256) NOT NULL,
    "description" text,
    "owner_id"    bigserial    NOT NULL,
    "updated_at"  timestamptz,
    "created_at"  timestamptz  NOT NULL DEFAULT (now())
);

CREATE TABLE "board_users"
(
    "board_id"   bigserial   NOT NULL,
    "user_id"    bigserial   NOT NULL,
    "role"       roles       NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "boards" ("owner_id");

CREATE UNIQUE INDEX ON "boards" ("owner_id", "name");

CREATE INDEX ON "board_users" ("board_id");

CREATE INDEX ON "board_users" ("user_id");

CREATE UNIQUE INDEX ON "board_users" ("board_id", "user_id");

-- ALTER TABLE "boards" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

-- ALTER TABLE "board_users" ADD FOREIGN KEY ("board_id") REFERENCES "boards" ("id");

-- ALTER TABLE "board_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- ALTER TABLE "wallets" ADD FOREIGN KEY ("board_id") REFERENCES "boards" ("id");

-- ALTER TABLE "transactions" ADD FOREIGN KEY ("board_id") REFERENCES "boards" ("id");

-- ALTER TABLE "notes" ADD FOREIGN KEY ("board_id") REFERENCES "boards" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
