#! /usr/bin/env sh

echo "creating schema ${ZK_AUTH_DB_NAME}"

PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "DROP SCHEMA IF EXISTS ${ZK_AUTH_DB_NAME} CASCADE;"
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "CREATE SCHEMA IF NOT EXISTS ${ZK_AUTH_DB_NAME} AUTHORIZATION ${ZK_AUTH_DB_USER};"
