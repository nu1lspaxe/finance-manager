#!/bin/bash

for i in {1..100}
do
  username="guest$i"
  email="example$i@gmail.com"
  curl -X POST http://localhost:8989/v1/users/signup \
       -H "Content-Type: application/json" \
       -d "{\"username\":\"$username\", \"email\":\"$email\"}" &
done

wait
