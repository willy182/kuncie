-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
	"id" serial PRIMARY KEY,
	"order_id" varchar(10) NOT NULL,
	"total" NUMERIC(10,2) NOT NULL,
	"created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamptz NULL DEFAULT NULL
);
CREATE UNIQUE INDEX index_orders ON orders USING btree ("order_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
