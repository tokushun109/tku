FROM golang:1.25-alpine

ENV ROOT=/api
WORKDIR ${ROOT}
COPY go.mod go.sum ./
RUN apk update && apk add --no-cache git && go mod download
COPY . .
