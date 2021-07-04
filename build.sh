#!/bin/bash


bin="zk-updater"

build_default() {
    go build -o dist/$bin cmd/zk-updater.go
}

build_cross () {
    GOOS=linux GOARCH=amd64 go build -o dist/$bin cmd/zk-updater.go
}

if [ "$1" == "cross" ]; then
    echo "building cross..."
    build_cross
else
    echo "building default..."
    build_default
fi

echo "done"
