#!/bin/bash -eu

go build -a -tags "netgo" -installsuffix netgo -ldflags="-s -w -extldflags \"-static\"" -o=./bin/wait-dockerd main.go
