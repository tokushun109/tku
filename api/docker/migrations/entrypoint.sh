#!/bin/sh

migrate -path /migrations -database "mysql://${DB_USER}:${DB_PASS}@tcp(${MYSQL_HOST})/${DB_NAME}" $@