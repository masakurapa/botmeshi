#!/bin/sh

rm -rf built
mkdir built go

export GO111MODULE=on
export GOOS=linux
export GOARCH=amd64

go build -o built/event event/event.go
go build -o built/interactive interactive/interactive.go
go build -o built/invoke invoke/invoke.go
