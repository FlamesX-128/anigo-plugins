#!/usr/bin/sh

VERSION='1.0.0'
NAME='consument'

go build -buildmode=plugin -o dist/linux-amd64/consument.so main.go
