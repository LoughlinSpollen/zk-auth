#! /usr/bin/env sh

source ./scripts/db-config.sh

echo "${ZK_AUTH_DB_NAME} database starting init - create tables'"

if psql ${ZK_AUTH_DB_NAME} -c '\q' 2>&1; then
    ./scripts/setup/f-db-create-auth-tables.sh
    echo "database ${ZK_AUTH_DB_NAME} database finished init - created auth tables'"
else
    echo "${ZK_AUTH_DB_NAME} database does not exist. Try running 'make db-setup'"
fi

