-- +goose Up
-- +goose StatementBegin
create table master_users
(
    uid          uuid default uuid_generate_v4() not null
        constraint uid_pk
            primary key,
    phone_number varchar(13)                     not null
        constraint phone_number_pk
            unique
);

comment on column master_users.phone_number is 'it should be unique';

alter table master_users
    owner to lucious;

create unique index master_users_phone_number_uindex
    on master_users (phone_number);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS master_users CASCADE;
DROP INDEX IF EXISTS master_users_phone_number_uindex;
-- +goose StatementEnd
