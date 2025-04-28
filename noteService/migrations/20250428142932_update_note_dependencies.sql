-- +goose Up
-- +goose StatementBegin
ALTER TABLE "notes" ADD FOREIGN KEY ("board_id") REFERENCES "boards" ("id") ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "notes"
DELETE CONSTRAINT IF EXISTS notes_board_id_fkey;
-- +goose StatementEnd
