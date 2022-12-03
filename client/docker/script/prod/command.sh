#!/bin/bash

set -e
yarn install --production=true
yarn build
yarn start