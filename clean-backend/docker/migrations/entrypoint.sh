#!/bin/sh

migrate -path /migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST})/${DB_NAME}" $@