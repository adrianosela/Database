#!/bin/bash +x

GOOS=linux GOARCH=amd64 go build -a -o Database-linux

docker build -t database.container .
