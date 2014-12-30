#!/bin/bash
docker run -e "GOPATH=/usr/src/asink/vendor" -e GOOS=openbsd -e GOARCH=amd64 --rm -v "$(pwd)":/usr/src/asink -w /usr/src/asink golang:1.4-cross make cross