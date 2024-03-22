#! /usr/bin/env sh

rm go.mod 2> /dev/null

go mod init zk_auth_service
go mod edit -replace zk_auth_service/lib/zk_auth_lib=../zk_auth_lib/go
go mod tidy
go build -v -o build/auth_service

echo "done"