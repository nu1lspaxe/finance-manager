CREATE TABLE IF NOT EXISTS "User" (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(20) NOT NULL,
    email VARCHAR(256) UNIQUE,
    phone_number NUMERIC(11,0),
    password VARCHAR(256),
    created_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_time TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "BankAccount" (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    account_number NUMERIC(15, 0) NOT NULL,
    bank_name VARCHAR(50),
    created_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_time TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "CreditCard" (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    card_number NUMERIC(16, 0) NOT NULL,
    cardholder_name VARCHAR(20),
    expiration_date DATE,
    created_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_time TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TYPE enum_record_type AS ENUM ('income', 'expense');
CREATE TYPE enum_record_source AS ENUM ('cash', 'credit_card', 'bank_account');

CREATE TABLE IF NOT EXISTS "Record" (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    amount NUMERIC(9,2) NOT NULL,
    transaction_date DATE NOT NULL,
    bank_account_id BIGINT,
    credit_card_id BIGINT,
    record_type enum_record_type NOT NULL,
    record_source enum_record_source NOT NULL,
    description TEXT,
    created_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_time TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_record_account_date ON "Record" (bank_account_id, transaction_date);
CREATE INDEX IF NOT EXISTS idx_record_card_date ON "Record" (credit_card_id, transaction_date);
CREATE INDEX IF NOT EXISTS idx_record_user_date ON "Record" (user_id, transaction_date);
