#!/usr/bin/sh

VERSION='1.0.0'
NAME='consument'

go build -buildmode=plugin -o "${NAME}@${VERSION}.so" main.go
