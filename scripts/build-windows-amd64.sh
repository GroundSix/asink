#!/bin/bash
docker run -e "GOPATH=/usr/src/asink/vendor" -e GOOS=windows -e GOARCH=amd64 -e EXT=.exe --rm -v "$(pwd)":/usr/src/asink -w /usr/src/asink golang:1.4-cross make cross