#! /usr/bin/env sh


echo "exporting environment"
rm .env 2> /dev/null
touch .env

cd database
./scripts/db-config.sh

cd ../grpc
./scripts/build.sh

cd ../zk_auth_service
./scripts/config.sh
rm go.mod 2> /dev/null
cd ..

echo "ZK_AUTH_DB_NAME=${ZK_AUTH_DB_NAME}" >> .env
echo "ZK_AUTH_DB_USER=${ZK_AUTH_DB_USER}" >> .env
echo "ZK_AUTH_DB_PASSWORD=${ZK_AUTH_DB_PASSWORD}" >> .env
echo "ZK_AUTH_DB_ADMIN=${ZK_AUTH_DB_ADMIN}" >> .env
echo "ZK_AUTH_DB_ADMIN_PASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD}" >> .env
echo "PGSSLMODE=${PGSSLMODE}" >> .env
echo "POSTGRES_DEBUG=${POSTGRES_DEBUG}" >> .env
echo "POSTGRES_DATA=${POSTGRES_DATA}" >> .env

echo "PGPORT=${PGPORT}" >> .env
echo "PGHOST=${PGHOST}" >> .env

echo "SERVICE_PORT=${SERVICE_PORT}" >> .env

echo "ZK_CPD_Q=${ZK_CPD_Q}" >> .env
echo "ZK_CPD_G=${ZK_CPD_Q}" >> .env
echo "ZK_CPD_A=${ZK_CPD_Q}" >> .env
echo "ZK_CPD_B=${ZK_CPD_Q}" >> .env

# leaving this here for sanity - not used
cat ./zk_client/.env >> .env

echo "building docker image"

docker volume prune -f  
docker-compose up

