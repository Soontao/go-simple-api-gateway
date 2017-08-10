#/bin/bash

export GOOS=linux GOARCH=amd64
go build -o "${PWD##*/}-${GOOS}-${GOARCH}"
export GOOS=linux GOARCH=386 
go build -o "${PWD##*/}-${GOOS}-${GOARCH}"
export GOOS=linux GOARCH=arm64 
go build -o "${PWD##*/}-${GOOS}-${GOARCH}"
export GOOS=darwin GOARCH=amd64
go build -o "${PWD##*/}-${GOOS}-${GOARCH}"
export GOOS=windows GOARCH=amd64
go build -o "${PWD##*/}-${GOOS}-${GOARCH}.exe"