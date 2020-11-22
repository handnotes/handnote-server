#!/bin/bash

. ./sendgrid.env

if [ "$1" = "swag" ]; then
  swag init
else
  docker-compose up -d
  swag init
  go run main.go
  docker-compose down
fi
