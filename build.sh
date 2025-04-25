#!/bin/bash

mkdir -vp ./builds

# Compilazione statica per Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./builds/proxy-linux-amd64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./builds/proxy-linux-arm64

# Compilazione statica per Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./builds/proxy-win-x64.exe
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ./builds/proxy-win-x86.exe

# Compilazione statica per macOS
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./builds/proxy-macos-darwin-amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ./builds/proxy-macos-darwin-arm64