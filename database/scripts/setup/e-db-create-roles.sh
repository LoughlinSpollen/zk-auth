#! /usr/bin/env sh

echo "creating role in schema ${ZK_AUTH_DB_NAME}"


# readwrite role for ZK_AUTH_DB_USER
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "create role readwrite;"
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "grant connect on database ${ZK_AUTH_DB_NAME} to readwrite;"
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "grant usage, create on schema ${ZK_AUTH_DB_NAME} to readwrite;"
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "grant select, insert, update, delete on all tables in schema ${ZK_AUTH_DB_NAME} to readwrite;"
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "alter default privileges in schema ${ZK_AUTH_DB_NAME} grant select, insert, update, delete on tables to readwrite;"
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "grant usage on all sequences in schema ${ZK_AUTH_DB_NAME} to readwrite;"
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "alter default privileges in schema ${ZK_AUTH_DB_NAME} grant usage on sequences to readwrite;"
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "alter default privileges grant all on functions to readwrite;"

PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "grant readwrite to ${ZK_AUTH_DB_USER};"
