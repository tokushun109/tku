#!/bin/bash

set -e

# login
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin ${ECR_REGISTRY}

# build
docker build -t tku_migration -f ./api/docker/migrations/Dockerfile ./api

# create tag
docker tag tku_migration:latest 418549683327.dkr.ecr.ap-northeast-1.amazonaws.com/tku_migration:latest

# push
docker push 418549683327.dkr.ecr.ap-northeast-1.amazonaws.com/tku_migration:latest
