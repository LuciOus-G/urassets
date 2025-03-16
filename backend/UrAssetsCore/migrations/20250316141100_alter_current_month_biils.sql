-- +goose Up
-- +goose StatementBegin
alter table current_month_income rename to current_month_bills;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table current_month_bills rename to current_month_income;
-- +goose StatementEnd
