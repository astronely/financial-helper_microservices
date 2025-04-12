-- +goose Up
-- +goose StatementBegin
CREATE TABLE "transactions"
(
    "id"         bigserial PRIMARY KEY,
    "owner_id"   bigserial   NOT NULL,
    "wallet_id"  bigserial   NOT NULL,
    "board_id"   bigserial   NOT NULL,
    "sum"        numeric     NOT NULL,
    "details_id" bigserial   NOT NULL,
    "updated_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transaction_details"
(
    "id"               bigserial PRIMARY KEY,
    "name"             varchar(128) NOT NULL,
    "category"         bigserial    NOT NULL,
    "transaction_date" date         NOT NULL
);

CREATE TABLE "transaction_categories"
(
    "id"          bigserial PRIMARY KEY,
    "name"        varchar(128) NOT NULL,
    "description" text         NOT NULL
);

CREATE INDEX ON "transactions" ("owner_id");

CREATE INDEX ON "transactions" ("board_id");

CREATE INDEX ON "transactions" ("wallet_id");

CREATE INDEX ON "transactions" ("owner_id", "wallet_id");

CREATE INDEX ON "transaction_details" ("category");

CREATE INDEX ON "transaction_details" ("transaction_date");

CREATE UNIQUE INDEX ON "transaction_categories" ("name");

ALTER TABLE "transactions"
    ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id")
    ON DELETE CASCADE;

ALTER TABLE "transactions"
    ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id")
    ON DELETE CASCADE;

-- ALTER TABLE "transactions" ADD FOREIGN KEY ("board_id") REFERENCES "boards" ("id"); // TODO: После добавления boards

ALTER TABLE "transactions"
    ADD FOREIGN KEY ("details_id") REFERENCES "transaction_details" ("id")
    ON DELETE CASCADE;

ALTER TABLE "transaction_details"
    ADD FOREIGN KEY ("category") REFERENCES "transaction_categories" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "transactions";
DROP TABLE IF EXISTS "transaction_details";
DROP TABLE IF EXISTS "transaction_categories";
-- +goose StatementEnd
