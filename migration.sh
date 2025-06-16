#!/bin/bash

source .env

sleep 2 && goose -dir "migrations" postgres "host=postgres port=$DB_PORT dbname=$DB_NAME user=$DB_USER password=$DB_PASSWORD" up -v