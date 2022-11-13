#!/usr/bin/sh

NAME='monochinos'

GOARCH='amd64'
GOOS='linux'

VERSION='1.0.1'
MODEL='1.0.0'

go build -buildmode=plugin -o dist/"$NAME@${VERSION}-v$MODEL-$GOOS-$GOARCH" main.go
