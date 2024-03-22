#! /usr/bin/env sh

echo "creating database ${ZK_AUTH_DB_NAME}"
createdb --template=template0 --locale=en_US.UTF-8 --encoding=UTF8 ${ZK_AUTH_DB_NAME}

psql -d ${ZK_AUTH_DB_NAME} -c "alter user ${ZK_AUTH_DB_ADMIN} with encrypted password '${ZK_AUTH_DB_ADMIN_PASSWORD}';"
psql -d ${ZK_AUTH_DB_NAME} -c "alter user ${ZK_AUTH_DB_ADMIN} with superuser;"
