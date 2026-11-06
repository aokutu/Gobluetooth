#!/bin/bash

# Define your local writable workspace path
export GOPATH="$HOME/go"
export PATH="$PATH:$GOPATH/bin"

# If no arguments are provided, print a usage hint
if [ $# -eq 0 ]; then
    echo "Usage: $0 [go command] [arguments...]"
    echo "Examples:"
    echo "  $0 get tinygo.org/x/bluetooth@v0.10.0"
    echo "  $0 run main.go"
    echo "  $0 mod tidy"
    exit 1
fi

# Pass all arguments safely straight to the go compiler wrapped in the local paths
go "$@"
