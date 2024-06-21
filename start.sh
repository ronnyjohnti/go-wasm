#! /bin/bash
go mod tidy
GOOS=js GOARCH=wasm go build -o web/main.wasm

# Inicializa servidor web na porta 5000
python3 -m http.server 5000 -d web/
# php -S 0.0.0.0:5000 -t web/
