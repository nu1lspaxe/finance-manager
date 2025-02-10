-- Table: User 

-- name: CreateUser :one
INSERT INTO "User" (username, email)
VALUES ($1, $2)
RETURNING *;

-- name: CheckUserExists :one
SELECT EXISTS (
  SELECT 1 FROM "User" WHERE id = $1
) AS user_exists;

-- name: CheckUserEmailExists :one
SELECT EXISTS (
  SELECT 1 FROM "User" WHERE email = $1
) AS email_exists;

-- name: GetUserById :one
SELECT * FROM "User" WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM "User" ORDER BY created_time;

-- name: UpdateUserEmail :exec
UPDATE "User" SET email = $2, updated_time = NOW() WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM "User" WHERE id = $1;

-- Table: BankAccount

-- name: CreateBankAccount :one
INSERT INTO "BankAccount" (user_id, account_number, bank_name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetBankAccountById :one
SELECT * FROM "BankAccount" WHERE id = $1;

-- name: ListBankAccounts :many
SELECT * FROM "BankAccount" ORDER BY created_time;

-- name: UpdateBankAccount :exec
UPDATE "BankAccount" SET account_number = $2, bank_name = $3, updated_time = NOW() WHERE id = $1;

-- name: DeleteBankAccount :exec
DELETE FROM "BankAccount" WHERE id = $1;

-- Table: CreditCard

-- name: CreateCreditCard :one
INSERT INTO "CreditCard" (user_id, card_number, cardholder_name, expiration_date)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCreditCardById :one
SELECT * FROM "CreditCard" WHERE id = $1;

-- name: ListCreditCards :many
SELECT * FROM "CreditCard" ORDER BY created_time;

-- name: UpdateCreditCard :exec
UPDATE "CreditCard" SET card_number = $2, cardholder_name = $3, expiration_date = $4, updated_time = NOW() WHERE id = $1;

-- name: DeleteCreditCard :exec
DELETE FROM "CreditCard" WHERE id = $1;

-- Table: Record

-- name: CreateRecord :one
INSERT INTO "Record" (user_id, amount, transaction_date, bank_account_id, credit_card_id, record_type, record_source, description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetRecordById :one
SELECT * FROM "Record" WHERE id = $1;

-- name: ListRecords :many
SELECT * FROM "Record" ORDER BY created_time;

-- name: UpdateRecord :exec
UPDATE "Record" SET amount = $2, transaction_date = $3, bank_account_id = $4, credit_card_id = $5, record_type = $6, record_source = $7, description = $8, updated_time = NOW() WHERE id = $1;

-- name: DeleteRecord :exec
DELETE FROM "Record" WHERE id = $1;

-- name: GetRecordsByUserAndDate :many
SELECT * FROM "Record"
WHERE user_id = $1
    AND transaction_date BETWEEN $2 AND $3;

-- name: GetRecordsByBankAccountAndDate :many
SELECT * FROM "Record"
WHERE bank_account_id = $1
  AND transaction_date BETWEEN $2 AND $3;

-- name: GetRecordsByCreditCardAndDate :many
SELECT * FROM "Record"
WHERE credit_card_id = $1
  AND transaction_date BETWEEN $2 AND $3;