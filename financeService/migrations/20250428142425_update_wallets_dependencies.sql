-- +goose Up
-- +goose StatementBegin
ALTER TABLE "wallets" ADD FOREIGN KEY ("board_id") REFERENCES "boards" ("id") ON DELETE CASCADE;
ALTER TABLE "transactions" ADD FOREIGN KEY ("board_id") REFERENCES "boards" ("id") ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "wallets"
DROP CONSTRAINT IF EXISTS wallets_board_id_fkey;

ALTER TABLE "transactions"
DROP CONSTRAINT IF EXISTS transactions_board_id_fkey;
-- +goose StatementEnd
