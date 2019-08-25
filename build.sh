#!/bin/sh

rm -rf built
mkdir built

export GO111MODULE=on
export GOOS=linux
export GOARCH=amd64

go build -o built/event event/event.go
go build -o built/interactive interactive/interactive.go
go build -o built/invoke-search invoke_search/invoke_search.go

cd built
zip event.zip event
zip interactive.zip interactive
zip invoke-search.zip invoke-search
