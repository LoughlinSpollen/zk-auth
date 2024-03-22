#! /usr/bin/env sh

. ./scripts/db-config.sh 
sleep 3

FILE=$POSTGRES_DATA/postmaster.pid
if [ ! -f "$FILE" ]; then
    echo "starting ${ZK_AUTH_DB_NAME} database"
    postgres -D $POSTGRES_DATA >logfile 2>&1 &
    sleep 2
else
    echo "${ZK_AUTH_DB_NAME} database already started"
fi
