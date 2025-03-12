-- +goose Up
-- +goose StatementBegin
alter table master_users
    add secure_pin varchar not null default '';

alter table master_users
    add email varchar
        constraint master_users_email_pk
            unique;

alter table master_users
    add password varchar;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table master_user drop column secure_pin;
alter table master_user drop column email;
-- +goose StatementEnd
