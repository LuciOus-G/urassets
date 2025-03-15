-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


-- Users Table
CREATE TABLE users (
                       id          UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                       full_name   VARCHAR(255) NOT NULL,
                       email       VARCHAR(255) NOT NULL UNIQUE,
                       password    VARCHAR(255) NOT NULL,
                       is_active   BOOLEAN DEFAULT TRUE,
                       created_at  TIMESTAMPTZ DEFAULT now(),
                       updated_at  TIMESTAMPTZ DEFAULT now(),
                       last_update TIMESTAMPTZ
);
COMMENT ON COLUMN users.id IS 'Auto Generate UUID V4';

-- Banks Table
CREATE TABLE banks (
                       id        UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                       bank_name VARCHAR(255) NOT NULL,
                       is_active BOOLEAN DEFAULT TRUE
);
COMMENT ON COLUMN banks.id IS 'Auto Generate UUID V4';

-- User Banks Table
CREATE TABLE user_banks (
                            id         UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                            bank_id    UUID NOT NULL REFERENCES banks (id),
                            user_id    UUID NOT NULL REFERENCES users (id),
                            created_at TIMESTAMPTZ DEFAULT now(),
                            updated_at TIMESTAMPTZ DEFAULT now(),
                            is_active  BOOLEAN DEFAULT TRUE
);
COMMENT ON COLUMN user_banks.id IS 'Auto Generate UUID V4';

-- Current Month Balance Table
CREATE TABLE current_month_balance (
                                       id               UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                                       bank_users_id    UUID NOT NULL REFERENCES user_banks (id),
                                       balance          NUMERIC NOT NULL,
                                       period_month     DATE,
                                       is_first_balance BOOLEAN DEFAULT FALSE
);
COMMENT ON COLUMN current_month_balance.id IS 'Auto Generate UUID V4';

-- Total Wealth Table
CREATE TABLE total_wealth (
                              id      UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                              user_id UUID NOT NULL REFERENCES users (id),
                              balance NUMERIC NOT NULL
);
COMMENT ON COLUMN total_wealth.id IS 'Auto Generate UUID V4';

-- Goals Table
CREATE TABLE goals (
                       id                UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                       user_id           UUID NOT NULL REFERENCES users (id),
                       goal_name         VARCHAR(255) NOT NULL,
                       goal_motivation   VARCHAR(255),
                       goal_descriptions TEXT NOT NULL,
                       goal_amount       NUMERIC NOT NULL,
                       priority          INTEGER DEFAULT 1
);
COMMENT ON COLUMN goals.id IS 'Auto Generate UUID V4';

-- Transactions Table
CREATE TABLE transactions (
                              id                   UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                              user_id              UUID NOT NULL REFERENCES users (id),
                              from_bank_id         UUID REFERENCES user_banks (id),
                              to_bank_id           UUID REFERENCES user_banks (id),
                              expenses_category_id UUID,
                              transfer_fee         NUMERIC,
                              amount               NUMERIC NOT NULL,
                              descriptions         TEXT
);
COMMENT ON COLUMN transactions.id IS 'Auto Generate UUID V4';

-- Current Month Income Table
CREATE TABLE current_month_income (
                                      id           UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                                      bank_user_id UUID NOT NULL REFERENCES user_banks (id),
                                      user_id      UUID NOT NULL REFERENCES users (id),
                                      amount       NUMERIC NOT NULL,
                                      period_month DATE NOT NULL,
                                      is_payed     BOOLEAN DEFAULT FALSE,
                                      name         VARCHAR(255) NOT NULL
);
COMMENT ON COLUMN current_month_income.id IS 'Auto Generate UUID V4';

-- User Income Category Table
CREATE TABLE user_income_category (
                                      id      UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                                      user_id UUID NOT NULL REFERENCES users (id),
                                      name    VARCHAR(255) NOT NULL
);

-- User Expenses Category Table
CREATE TABLE user_expenses_category (
                                        id      UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                                        user_id UUID NOT NULL REFERENCES users (id),
                                        name    VARCHAR(255) NOT NULL
);

-- Default Income Category Table
CREATE TABLE default_income_category (
                                         id   UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                                         name VARCHAR(255) NOT NULL
);

-- Default Expenses Category Table
CREATE TABLE default_expenses_category (
                                           id   UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                                           name VARCHAR(255) NOT NULL
);

-- Crypto Assets Total Table
CREATE TABLE crypto_assets_total (
                                     id      UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                                     user_id UUID NOT NULL REFERENCES users (id),
                                     name    VARCHAR(255) NOT NULL,
                                     symbol  VARCHAR(255) NOT NULL,
                                     value   NUMERIC NOT NULL
);

-- Crypto Assets History Table
CREATE TABLE crypto_assets_history (
                                       id                     UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                                       user_id                UUID NOT NULL REFERENCES users (id),
                                       crypto_assets_total_id UUID REFERENCES crypto_assets_total (id),
                                       name                   VARCHAR(255) NOT NULL,
                                       symbol                 VARCHAR(255) NOT NULL,
                                       value                  NUMERIC NOT NULL,
                                       created_at             TIMESTAMPTZ DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS banks CASCADE;
DROP TABLE IF EXISTS user_banks CASCADE;
DROP TABLE IF EXISTS current_month_balance CASCADE;
DROP TABLE IF EXISTS total_wealth CASCADE;
DROP TABLE IF EXISTS goals CASCADE;
DROP TABLE IF EXISTS transactions CASCADE;
DROP TABLE IF EXISTS current_month_income CASCADE;
DROP TABLE IF EXISTS user_income_category CASCADE;
DROP TABLE IF EXISTS user_expenses_category CASCADE;
DROP TABLE IF EXISTS default_income_category CASCADE;
DROP TABLE IF EXISTS default_expenses_category CASCADE;
DROP TABLE IF EXISTS crypto_assets_total CASCADE;
DROP TABLE IF EXISTS crypto_assets_history CASCADE;
-- +goose StatementEnd
