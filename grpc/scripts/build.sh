#! /usr/bin/env sh

echo "GPRC/Protobuf generation"

echo "Generateing Python bindings for zk auth client"
# copy source to temp
mkdir -p tmp/build/protos/zk_auth
cp zk-auth-api.proto tmp/build/protos/zk_auth

# Golang naming convention is different to Python.
# Golang uses camelCase or lowercase with underscores in file names.
# Python uses camelCase and Python doesn't accept underscores in package names.
# A temp build directory is used as a workaround.
cd tmp

python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. build/protos/zk_auth/zk-auth-api.proto
rm -rf ../../zk_client/build/protos/zk_auth
mkdir -p ../../zk_client/build/protos/zk_auth
cp build/protos/zk_auth/*.py ../../zk_client/build/protos/zk_auth

# generate golang bindings
echo "Generating Golang bindings zk auth service"
mkdir -p ../../zk_auth_service/build/protos/zk_auth
protoc --proto_path=build/protos/zk_auth/ \
    --go_out=../../zk_auth_service/build/protos/zk_auth \
    --go-grpc_out=../../zk_auth_service/build/protos/zk_auth \
    --go_opt=Mzk-auth-api.proto=./ \
    --go-grpc_opt=Mzk-auth-api.proto=./ \
    build/protos/zk_auth/zk-auth-api.proto  

cd ..
rm -rf tmp

