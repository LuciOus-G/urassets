-- +goose Up
-- +goose StatementBegin
create table master_wallet
(
    wid            uuid    default uuid_generate_v4() not null
        constraint master_wallet_pk
            primary key,
    master_user_id uuid
        constraint master_wallet_master_users_uid_fk
            references master_users
            on update cascade on delete cascade,
    balance        varchar default '0'                not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS master_wallet CASCADE;
-- +goose StatementEnd
