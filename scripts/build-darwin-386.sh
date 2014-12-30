#!/bin/bash
docker run -e "GOPATH=/usr/src/asink/vendor" -e GOOS=darwin -e GOARCH=386 --rm -v "$(pwd)":/usr/src/asink -w /usr/src/asink golang:1.4-cross make cross