# ZKP Challenge

## Overview

This is a zk authenication service and client spike.

## Install

* This will install a `golang` `postgres` `python` `protobufs` stack on your mac-os dev machine. It assumes `docker` is installed too.

* Install from `./install-dev`

## Run

* Run from `make run` or `make test` within the top level directories.
* The server `make test` outputs some coverage metrics. It's about 80% which is fine. Testing is a law of demishing returns.

## Background

* This spike adapts the `Chaum–Pedersen Protocol` to support 1-factor authentication, that is, the exact matching of a number (registration password) stored during registration and another number (login password) generated during the login process.

### Layout

* The top level directorys are isolated and can eventually graduate to repos. This includes a Python client, a Golang server, a Posgres DB and a gPRC API.

## Solution

* This implements ZK Proof auth client, server and persistnece system.
* The `make` commands work.
* Run the `make run-servers` from the root.
* The docker images are buiding. The `docker-compose` works too. Check out the associated `make` commands to ease automation: `make compose-up`and `make compose-down`. It's probably best to use these commands to ensure volumes are pruned and the shared lib is included in the build.
