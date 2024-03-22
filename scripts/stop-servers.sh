#! /usr/bin/env sh

sleep 1s
cd zk_auth_service/
. ./scripts/config.sh
cd ..
lsof -ti:$SERVICE_PORT | xargs kill
sleep 1s

echo "stopping database"
cd database
./scripts/db-stop.sh
cd ..