#! /bin/bash

go mod tidy
GOOS=js GOARCH=wasm go build -o web/main.wasm

# Copy the Go WebAssembly initializing script to application
cp $(go env GOROOT)/misc/wasm/wasm_exec.js web/

# Starts a web server on port 5000
python3 -m http.server 5000 -d web/
# php -S 0.0.0.0:5000 -t web/
