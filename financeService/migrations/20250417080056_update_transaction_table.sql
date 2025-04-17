-- +goose Up
-- +goose StatementBegin
-- 1. Сначала создаём ENUM, если его ещё нет
CREATE TYPE "transaction_type" AS ENUM ('expense', 'income', 'transfer');

-- 2. Добавляем недостающие столбцы
ALTER TABLE "transactions"
    ADD COLUMN IF NOT EXISTS "from_wallet_id" bigint,
    ADD COLUMN IF NOT EXISTS "to_wallet_id" bigint DEFAULT NULL,
    ADD COLUMN IF NOT EXISTS "type" transaction_type NOT NULL DEFAULT 'expense';

UPDATE "transactions"
    SET "from_wallet_id" = "wallet_id";
-- 3. Удаляем/переименовываем старые поля, если они больше не нужны
-- Пример: переименовать "sum" → "amount", если поле уже было
ALTER TABLE "transactions"
    RENAME COLUMN "sum" TO "amount";

-- 4. Добавляем/обновляем связи и индексы
CREATE INDEX IF NOT EXISTS transactions_from_wallet_id_idx ON "transactions" ("from_wallet_id");
CREATE INDEX IF NOT EXISTS transactions_to_wallet_id_idx ON "transactions" ("to_wallet_id");
CREATE INDEX IF NOT EXISTS transactions_owner_id_from_wallet_id_idx ON "transactions" ("owner_id", "from_wallet_id");

ALTER TABLE "transactions"
    ADD FOREIGN KEY ("to_wallet_id") REFERENCES "wallets" ("id") ON DELETE SET NULL;

ALTER TABLE "transactions"
    ADD FOREIGN KEY ("from_wallet_id") REFERENCES "wallets" ("id") ON DELETE CASCADE;

ALTER TABLE "transactions"
    ALTER COLUMN "from_wallet_id" SET NOT NULL;

-- Удаление внешнего ключа
ALTER TABLE "transactions"
    DROP CONSTRAINT IF EXISTS transactions_wallet_id_fkey;
-- Удаление индексов wallet_id
DROP INDEX IF EXISTS transactions_wallet_id_idx;
DROP INDEX IF EXISTS transactions_owner_id_wallet_id_idx;

-- Удаление столбца wallet_id
ALTER TABLE "transactions"
DROP COLUMN IF EXISTS "wallet_id";

-- Удаление каскадного удаления предыдущей миграции
ALTER TABLE "transaction_details"
    DROP CONSTRAINT IF EXISTS transaction_details_category_fkey;

ALTER TABLE "transaction_details"
    ADD FOREIGN KEY ("category") REFERENCES "transaction_categories" ("id");

ALTER TABLE "wallets"
    ADD CONSTRAINT balance_positive CHECK ("balance" > 0);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- 1. Удаляем CHECK constraint с wallets
ALTER TABLE "wallets"
DROP CONSTRAINT IF EXISTS balance_positive;

-- 2. Удаляем внешний ключ с transaction_details
ALTER TABLE "transaction_details"
DROP CONSTRAINT IF EXISTS transaction_details_category_fkey;

-- 3. Восстанавливаем старый внешний ключ с каскадным удалением
ALTER TABLE "transaction_details"
    ADD CONSTRAINT transaction_details_category_fkey
        FOREIGN KEY ("category") REFERENCES "transaction_categories" ("id") ON DELETE CASCADE;

-- 4. Восстанавливаем столбец wallet_id
ALTER TABLE "transactions"
    ADD COLUMN IF NOT EXISTS "wallet_id" bigint;

-- 5. Переносим данные из from_wallet_id обратно в wallet_id
UPDATE "transactions"
SET "wallet_id" = "from_wallet_id";

-- 6. Добавляем NOT NULL для wallet_id
ALTER TABLE "transactions"
    ALTER COLUMN "wallet_id" SET NOT NULL;

-- 7. Восстанавливаем внешний ключ
ALTER TABLE "transactions"
    ADD CONSTRAINT transactions_wallet_id_fkey
        FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id") ON DELETE CASCADE;

-- 8. Восстанавливаем индексы
CREATE INDEX IF NOT EXISTS transactions_wallet_id_idx ON "transactions" ("wallet_id");
CREATE INDEX IF NOT EXISTS transactions_owner_id_wallet_id_idx ON "transactions" ("owner_id", "wallet_id");

-- 9. Удаляем новые внешние ключи
ALTER TABLE "transactions"
DROP CONSTRAINT IF EXISTS transactions_from_wallet_id_fkey;

ALTER TABLE "transactions"
DROP CONSTRAINT IF EXISTS transactions_to_wallet_id_fkey;

-- 10. Удаляем новые индексы
DROP INDEX IF EXISTS transactions_from_wallet_id_idx;
DROP INDEX IF EXISTS transactions_to_wallet_id_idx;
DROP INDEX IF EXISTS transactions_owner_id_from_wallet_id_idx;

-- 11. Удаляем новые столбцы
ALTER TABLE "transactions"
DROP COLUMN IF EXISTS "from_wallet_id",
    DROP COLUMN IF EXISTS "to_wallet_id",
    DROP COLUMN IF EXISTS "type";

-- 12. Переименовываем amount обратно в sum
ALTER TABLE "transactions"
    RENAME COLUMN "amount" TO "sum";

-- 13. Удаляем ENUM тип
DROP TYPE IF EXISTS "transaction_type";
-- +goose StatementEnd

