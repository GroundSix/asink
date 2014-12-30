#!/bin/bash
docker run -e "GOPATH=/usr/src/asink/vendor" --rm -v "$(pwd)":/usr/src/asink -w /usr/src/asink asink make