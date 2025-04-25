#!/bin/bash

mkdir -vp ./builds

GOOS=linux GOARCH=amd64 go build -o ./builds/proxy-linux-amd64
GOOS=linux GOARCH=arm64 go build -o ./builds/proxy-linux-arm64
GOOS=windows GOARCH=amd64 go build -o ./builds/proxy-win-x64.exe
GOOS=darwin GOARCH=amd64 go build -o ./builds/proxy-macos-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -o ./builds/proxy-macos-darwin-arm64
