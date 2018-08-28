#!/bin/bash
GOPATH=$(pwd) go build -o build/main main && build/main