#!/bin/bash

token=$(curl -s -d '{"username": "testname", "password": "pass"}' -X POST "$1:8000/create")
username=$(curl -s -H "Content-Type: application/json" -d '{"token": "'"$token"'"}' -X POST "$1:5000/verifier")

if [ "$username" = "testname" ]; then
  echo "GOOD!"
else
  echo "BAD!"
fi
