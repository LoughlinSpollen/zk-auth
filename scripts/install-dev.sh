#! /usr/bin/env zsh

echo "installing gRPC"
cd grpc
./scripts/install-dev.sh

echo "installing database"
cd ../database
./scripts/install-dev.sh

echo "installing python dependencies"
cd ../zk_client/
./scripts/install-dev.sh

echo "installing go-lang environment"
cd ../zk_auth_service
./scripts/install-dev.sh

cd ..

