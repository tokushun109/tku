#!/bin/sh

migrate -path /migrations -database "mysql://${DB_USER}:${DB_PASS}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${DB_NAME}" $@
