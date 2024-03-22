#! /usr/bin/env sh

echo "building protobuffers"
cd grpc
./scripts/build.sh
cd ..

echo "building all services"

cd zk_auth_service
rm go.mod 2> /dev/null
cd ..

echo "building zk-auth-service"
docker build -t zk-auth-service:v1 -f zk-auth-service.Dockerfile .

echo "building zk-client"
docker build -t zk-client:v1 -f zk-client.Dockerfile .
