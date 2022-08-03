#!/usr/bin/env bash

BASE_PATH=./
OUT_PATH=./bin

# shellcheck disable=SC2231
export GO111MODULE=on
env GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o $OUT_PATH/main $BASE_PATH/"$handler"/lambda.go
cd $OUT_PATH
zip main.zip main
