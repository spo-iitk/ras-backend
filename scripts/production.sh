#!/usr/bin/env bash

# Pull latest commit
# git pull origin main --force

# # Install `things`
# go get -d -v ./...
# go install -v ./...

service nginx start
go build -o server ./cmd
# ./server
air
