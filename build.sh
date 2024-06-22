#!/usr/bin/bash

export GOLANG_VERSION="1.22.3"

go mod tidy -e

go build -ldflags "-w -s -X 'main.Version=0.9.0'" -o ./metadata ./cli
