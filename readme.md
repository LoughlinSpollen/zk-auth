# ZKP Challenge

## Overview

This is a zk authenication service and client spike.

## Install

* This will install a `golang` `postgres` `python` `protobufs` stack on your mac-os dev machine. It assumes `docker` is installed too.

* Install from `./install-dev`

## Run

* Run from `make run`

## Background

* As a Spike that will grow into a product. It is a mono-repo with with isolated services that will eventually graduate to their own repo, so it's isolated in this fashion but with a common architecture. The common pattern across client and services hopefully makes it easier to follow.

* This spike adapts the `Chaum–Pedersen Protocol` to support 1-factor authentication, that is, the exact matching of a number (registration password) stored during registration and another number (login password) generated during the login process.

### Layout

* The top level directorys are isolated and can eventually graduate to repos. This includes a Python client, a Golang server, a Posgres DB and a gPRC API.

## Solution

* This implements ZK Proof auth client, server and persistnece system.
* The `make` commands work.
* Run the `make run-servers` from the root.
* The docker images build, I haven't had a chance to check if the `docker-compose` actually works.