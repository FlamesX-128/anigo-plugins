#!/usr/bin/sh

VERSION='1.0.0'
NAME='monochinos'

go build -buildmode=plugin -o "${NAME}@${VERSION}.so" main.go
