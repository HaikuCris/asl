#!/bin/bash

ASL_PROJECT=$(cd `dirname $0`;pwd)

cd $ASL_PROJECT
mkdir -p $ASL_PROJECT/out
rm -rf $ASL_PROJECT/out/*
# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $ASL_PROJECT/out/aslgo_linux_arm64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $ASL_PROJECT/out/aslgo_linux_amd64
# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o $ASL_PROJECT/out/aslgo_windows_arm64.exe
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $ASL_PROJECT/out/aslgo_windows_amd64.exe
# Mac
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $ASL_PROJECT/out/aslgo_darwin_arm64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $ASL_PROJECT/out/aslgo_darwin_amd64
