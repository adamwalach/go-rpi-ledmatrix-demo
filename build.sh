#!/bin/bash

go version
CC_PATH=/usr/bin/arm-linux-gnueabi-gcc
CXX_PATH=/usr/bin/arm-linux-gnueabi-cpp

CC=$CC_PATH CXX=$CXX_PATH GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=1 go build -o=main *.go
