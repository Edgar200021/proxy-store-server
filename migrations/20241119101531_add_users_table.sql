-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id TEXT PRIMARY KEY NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
