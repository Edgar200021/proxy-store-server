-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id TEXT PRIMARY KEY NOT NULL UNIQUE,
	balance DECIMAL(12,2) NOT NULL DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
