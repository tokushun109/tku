#!/bin/sh

migrate -path /migrations -database "mysql://${DB_USER}:${DB_PASS}@${PROTOCOL}/${DB_NAME}" $@