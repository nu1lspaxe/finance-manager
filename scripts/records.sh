#!/bin/bash

generate_random_amount() {
  echo "scale=2; 1 + $RANDOM / 32767 * 1000" | bc
}

generate_random_date() {
  year=$((RANDOM % 5 + 2020)) 
  month=$((RANDOM % 12 + 1))   
  day=$((RANDOM % 28 + 1))   
  printf "%04d-%02d-%02d\n" $year $month $day
}

for user_id in {1..100}
do
  for year in {2022..2024}
  do
    for month in {1..12}
    do
      for i in {1..3}
      do
        amount=$(generate_random_amount)
        transaction_date=$(printf "%04d-%02d-%02d" $year $month $((RANDOM % 28 + 1)))
        curl -X POST http://localhost:8989/v1/records/create \
             -H "Content-Type: application/json" \
             -d "{\"user_id\":$user_id, \"amount\":$amount, \"transaction_date\":\"$transaction_date\", \"record_type\":\"income\", \"record_source\":\"cash\", \"description\":\"\"}" &
      done
      for i in {1..3}
      do
        amount=$(generate_random_amount)
        transaction_date=$(printf "%04d-%02d-%02d" $year $month $((RANDOM % 28 + 1)))
        curl -X POST http://localhost:8989/v1/records/create \
             -H "Content-Type: application/json" \
             -d "{\"user_id\":$user_id, \"amount\":$amount, \"transaction_date\":\"$transaction_date\", \"record_type\":\"expense\", \"record_source\":\"cash\", \"description\":\"\"}" &
      done
    done
  done
done

wait
