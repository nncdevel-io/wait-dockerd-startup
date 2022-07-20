#!/bin/bash -eu

go build -a -tags "netgo" -installsuffix netgo -ldflags="-s -w -extldflags \"-static\"" -o=./dist/wait-dockerd-startup main.go
