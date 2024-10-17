FROM --platform=linux/amd64 golang:1.21-alpine

ENV ROOT=/api
WORKDIR ${ROOT}
COPY ./ ${ROOT}

RUN apk update && apk add git && go mod download
