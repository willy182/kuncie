-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items (
	"id" serial PRIMARY KEY,
	"order_id" varchar(10) NOT NULL,
	"sku" varchar(6) NOT NULL,
	"qty" INTEGER NOT NULL,
	"amount" NUMERIC(8,2) NOT NULL,
	"created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamptz NULL DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
