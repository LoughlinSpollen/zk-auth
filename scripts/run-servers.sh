#! /usr/bin/env sh

pwd
cd zk_auth_service/
. ./scripts/config.sh
lsof -ti:$SERVICE_PORT | xargs kill
cd ..


echo "building protobuffers"
cd grpc
./scripts/build.sh
cd ..

echo "starting database"
cd database
./scripts/db-stop.sh
./scripts/db-delete.sh
sleep 3
./scripts/db-start.sh
./scripts/db-setup.sh
./scripts/db-init.sh
cd ..

echo "starting server"
cd zk_auth_service/
./scripts/build.sh
./build/auth_service &
cd ..


echo "starting client"
cd zk_client/
python main.py &
cd ..


sleep 3

