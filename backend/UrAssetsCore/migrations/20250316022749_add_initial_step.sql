-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_steps (
    user_id   UUID NOT NULL REFERENCES users (id) PRIMARY KEY,
    step_1    BOOL default false,
    step_2    BOOL default false,
    step_3    BOOL default false,
    is_complete BOOL default false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_steps CASCADE;
-- +goose StatementEnd
