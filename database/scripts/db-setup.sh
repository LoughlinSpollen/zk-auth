#! /usr/bin/env sh

./scripts/db-config.sh

echo "database ${ZK_AUTH_DB_NAME} setup starting"
sleep 2s # Waits 2 seconds.

./scripts/setup/a-db-create-admin.sh
./scripts/setup/b-db-create-database.sh

./scripts/setup/c-db-create-user.sh
./scripts/setup/d-db-create-schema.sh
./scripts/setup/e-db-create-roles.sh

echo "database ${ZK_AUTH_DB_NAME} setup complete"
