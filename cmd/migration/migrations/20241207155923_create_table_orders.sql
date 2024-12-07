-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
	"id" uuid DEFAULT gen_random_uuid(),
	"type" VARCHAR(20) NOT NULL,
	"transaction_type" VARCHAR(20) NOT NULL,
	"status" VARCHAR(20) NOT NULL,
	"amount" DECIMAL(10,2) NOT NULL,
	"balance_before" DECIMAL(10,2),
	"balance_after" DECIMAL(10,2),
	"remarks" VARCHAR(100),
	"user_id" uuid NOT NULL,
	"transfer_to_user_id" uuid,
	"created_at" TIMESTAMPTZ(6),
	"updated_at" TIMESTAMPTZ(6),
	PRIMARY KEY ("id"),
	CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
	CONSTRAINT fk_transfer_user FOREIGN KEY(transfer_to_user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
