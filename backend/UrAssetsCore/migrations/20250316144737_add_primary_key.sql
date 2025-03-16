-- +goose Up
-- +goose StatementBegin
alter table current_month_income
    add id uuid default uuid_generate_v4() not null
        constraint current_month_income_pk
            primary key;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
