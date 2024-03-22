
#! /usr/bin/env sh

source ./scripts/db-config.sh

# CASCADE is used to remove all data from related tables but not used - it is not necessary. Here for completeness.
echo "removing all data from tables in ${ZK_AUTH_DB_NAME}"
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "TRUNCATE ${ZK_AUTH_DB_NAME}.auth CASCADE;"
