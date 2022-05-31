-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
	"id" serial PRIMARY KEY,
	"sku" varchar(6) NOT NULL,
	"name" varchar(255) NOT NULL,
	"price" NUMERIC(7,2) NOT NULL,
	"stock" INTEGER NOT NULL,
	"created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamptz NULL DEFAULT NULL
);
CREATE UNIQUE INDEX index_products ON products USING btree ("sku");

INSERT INTO products (sku, name, price, stock, created_at, updated_at)
VALUES
    ('120P90', 'Google Home', 49.99, 10, now(), null),
    ('43N23P', 'MacBook Pro', 5399.99, 5, now(), null),
    ('A304SD', 'Alexa Speaker', 109.50, 10, now(), null),
    ('234234', 'Raspberry Pi B', 30.00, 2, now(), null);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
