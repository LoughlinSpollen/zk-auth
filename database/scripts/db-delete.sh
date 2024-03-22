#! /usr/bin/env sh

./scripts/db-config.sh

rm -rf $POSTGRES_DATA
mkdir -p $POSTGRES_DATA

echo "the database data location is $POSTGRES_DATA, local to the development environment."
initdb -D $POSTGRES_DATA

echo "database ${ZK_AUTH_DB_NAME} deleted"
