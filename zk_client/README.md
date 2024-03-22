# ZKP Client

## Overview

* Simple client written in python 3.7, wires up the gRPC bindings and zk_auth_lib.

## Install

* Install from `make install-dev` to checks if python is installed and adds packages.

## Run

* Run from `make run`. Ensure the server is running and the gPRC protobuf binds are built. 

## Test

* Run from `make run` to run an e2e test, once an instance of the server is running.
* There are no unit tests as there's nothing to test.

## Solution

* This client wires up the shared ZK auth python package to the gRPC bindings.
* It creates a password and registers it with the server. Then it authenticates and exits.
* To ease with deployment to a cloud, it adds a back-off/retry to gRPC.
