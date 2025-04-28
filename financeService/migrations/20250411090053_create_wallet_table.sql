-- +goose Up
-- +goose StatementBegin
CREATE TABLE "wallets"
(
    "id"         bigserial PRIMARY KEY,
    "owner_id"   bigserial    NOT NULL,
    "board_id"   bigserial    NOT NULL,
    "name"       varchar(128) NOT NULL,
    "balance"    numeric(18, 2)        DEFAULT 0,
    "updated_at" timestamptz,
    "created_at" timestamptz  NOT NULL DEFAULT (now())
);
CREATE INDEX ON "wallets" ("owner_id");
CREATE INDEX ON "wallets" ("board_id");
CREATE UNIQUE INDEX ON "wallets" ("owner_id", "board_id", "name");

ALTER TABLE "wallets" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "wallets";
-- +goose StatementEnd
