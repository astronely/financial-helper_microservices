-- +goose Up
-- +goose StatementBegin
CREATE TABLE "notes"
(
    "id"              bigserial PRIMARY KEY,
    "board_id"        bigint   NOT NULL,
    "owner_id"        bigint   NOT NULL,
    "content"         text        NOT NULL,
    "status"          bool        NOT NULL DEFAULT false,
    "performer_id"    bigint,
    "completion_date" timestamptz,
    "updated_at"      timestamptz,
    "created_at"      timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "notes" ("owner_id");

CREATE INDEX ON "notes" ("performer_id");

CREATE INDEX ON "notes" ("status");

CREATE INDEX ON "notes" ("owner_id", "status");

CREATE INDEX ON "notes" ("performer_id", "status");

-- ALTER TABLE "notes" ADD FOREIGN KEY ("board_id") REFERENCES "boards" ("id");

ALTER TABLE "notes" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "notes" ADD FOREIGN KEY ("performer_id") REFERENCES "users" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "notes";
-- +goose StatementEnd
