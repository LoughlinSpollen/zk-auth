#! /usr/bin/env sh

echo "test"
# setup the database
. ../database/scripts/db-config.sh

echo "starting database"
cd ../database
./scripts/db-stop.sh
./scripts/db-delete.sh
./scripts/db-start.sh
./scripts/db-setup.sh
./scripts/db-init.sh
./scripts/db-truncate.sh

echo "starting server"
cd ../zk_auth_service
./scripts/build.sh
./build/auth_service &
cd ..

echo "Unit and integration tests"
rm -rf ../test-results
mkdir ../test-results
ginkgo -timeout=60s -r -v --output-dir=../test-results


cd ../database
./scripts/db-stop.sh
cd ..

lsof -ti:$SERVICE_PORT | xargs kill
echo "Tests complete"

