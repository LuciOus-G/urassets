-- +goose Up
-- +goose StatementBegin
create table current_month_income
(
    bank_user_id uuid    not null
        constraint current_month_income_user_banks_id_fk
            references user_banks,
    user_id      uuid    not null
        constraint current_month_income_users_id_fk
            references users,
    amount       numeric not null,
    period_month date,
    type         varchar(255),
    name         varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists current_month_income;
-- +goose StatementEnd
