#! /usr/bin/env sh
./scripts/db-config.sh

FILE=$POSTGRES_DATA/postmaster.pid
echo "Database location: " $FILE
if test -f "$FILE"; then
    echo "${ZK_AUTH_DB_NAME} stopping database"
    pg_ctl -D $POSTGRES_DATA stop 2>&1 &
else
    echo "${ZK_AUTH_DB_NAME} database was not started"
fi
