#! /usr/bin/env zsh
echo "Installing grpc core and protobufs"

# python
pip install grpcio-tools

# go
echo
brew tap grpc/grpc
brew install grpc
brew install protoc-gen-go
