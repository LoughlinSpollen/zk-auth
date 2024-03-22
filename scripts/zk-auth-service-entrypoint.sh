#! /usr/bin/env sh
echo "Waiting for ${ZK_AUTH_DB_NAME} database"
RETRIES=5

IS_DB_READY="pg_isready -h ${PGHOST} -U ${ZK_AUTH_DB_USER} -d ${ZK_AUTH_DB_NAME}"

eval $IS_DB_READY > /dev/null 2>&1
until [ $? -eq 0 ];
do
    RETRIES=$(( RETRIES - 1 ))
    if [ $RETRIES -eq 0 ] ; then
        echo "Failed to find database ${ZK_AUTH_DB_NAME}, bye!"
        exit
    fi

    echo "Waiting for postgres server to start, ${RETRIES} remaining attempts..."
    sleep 5s
    $IS_DB_READY > /dev/null 2>&1
done

/go/bin/zk_auth_service