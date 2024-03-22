#! /usr/bin/env sh

cd ../database
. ./scripts/db-config.sh
cd ../zk_auth_service
export PGPORT=5432
export PGHOST="0.0.0.0"

export SERVICE_PORT=1025
export ZK_CPD_Q=10003
export ZK_CPD_G=3
export ZK_CPD_A=10
export ZK_CPD_B=13
