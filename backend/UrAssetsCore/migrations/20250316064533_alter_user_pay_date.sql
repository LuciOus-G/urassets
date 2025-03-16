-- +goose Up
-- +goose StatementBegin
alter table users add pay_date timestamptz;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table users drop column pay_date;
-- +goose StatementEnd
