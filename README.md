# finance-manager

## Dev Guide

1. Set up `.env` with template `.env-template`
   
2. Run `docker-compose`

   ```bash
   docker compose --env-file=.env up --build -d
   ```
3. Run server

   ```bash
   go run cmd/main.go
   ```
4. Manual Test (with shell script)
   
   ```bash
   ./scripts/users.sh   # Create 100 users
   ./scripts/records.sh # Create records for 100 users
   ```

## Check List

- [X] `/v1/users/signup`
- [X] `/v1/records/create`

## Test script

1. Create user

- Single
   ```bash
   curl -X POST http://localhost:8989/v1/users/signup \
      -H "Content-Type: application/json" \
      -d '{"username":"guest1", "email":"example1@gmail.com"}'
   ```

- Multiple
  ```bash
  ./scripts/users.sh 
  ```


2. Add record

- Single
   ```bash
   curl -X POST http://localhost:8989/v1/records/create \
      -H "Content-Type: application/json" \
      -d '{
            "user_id": 1,
            "amount": 100.00,
            "transaction_date": "2025-02-10",
            "record_type": "expense",
            "record_source": "cash",
            "description": "Grocery shopping"
            }'
   ```

- Multiple
  ```bash
  ./scripts/records.sh
  ```