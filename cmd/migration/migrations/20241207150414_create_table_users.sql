-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	"id" uuid DEFAULT gen_random_uuid(),
	"first_name" VARCHAR(100) NOT NULL,
	"last_name" VARCHAR(100),
	"phone_number" VARCHAR(16) NOT NULL,
	"pin" VARCHAR(255) NOT NULL,
	"salt" VARCHAR(255),
	"saldo" DECIMAL(10,2),
	"address" TEXT,
	"created_at" TIMESTAMPTZ(6),
	"updated_at" TIMESTAMPTZ(6),
	PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_phone_number" ON users USING btree ("phone_number");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "idx_phone_number";
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
