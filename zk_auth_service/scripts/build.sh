#! /usr/bin/env sh

go mod edit -replace zk_auth_service/lib/zk_auth_lib=../zk_auth_lib/go
go mod tidy
go build -v -o build/auth_service

echo "done"