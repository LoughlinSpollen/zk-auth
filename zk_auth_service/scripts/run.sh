#! /usr/bin/env sh
./scripts/config.sh

go build -v -o ./build/auth_service -tags debug
./build/auth_service