-- +goose Up
-- +goose StatementBegin
alter table current_month_balance rename to current_month_wealth;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table current_month_wealth rename to current_month_balance;
-- +goose StatementEnd
