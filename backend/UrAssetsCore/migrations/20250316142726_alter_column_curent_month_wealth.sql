-- +goose Up
-- +goose StatementBegin
alter table current_month_wealth
    rename column bank_users_id to users_id;

alter table current_month_wealth
drop column is_first_balance;

alter table current_month_wealth
drop constraint current_month_balance_bank_users_id_fkey;

alter table current_month_wealth
    add constraint current_month_balance_bank_users_id_fkey
        foreign key (users_id) references users;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table current_month_wealth
    rename column users_id to bank_users_id;

alter table current_month_wealth
add column is_first_balance;

alter table current_month_wealth
add constraint current_month_balance_bank_users_id_fkey
    foreign key (bank_users_id) references user_banks;

alter table current_month_wealth
    drop constraint current_month_balance_bank_users_id_fkey;
-- +goose StatementEnd
