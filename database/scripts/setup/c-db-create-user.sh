#! /usr/bin/env sh

echo "creating user is ${ZK_AUTH_DB_USER}"
createuser ${ZK_AUTH_DB_USER}

PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "alter user ${ZK_AUTH_DB_USER} with encrypted password '${ZK_AUTH_DB_PASSWORD}';"
