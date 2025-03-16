-- +goose Up
-- +goose StatementBegin
alter table user_banks
    add balance numeric;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table user_banks
    drop balance;
-- +goose StatementEnd
