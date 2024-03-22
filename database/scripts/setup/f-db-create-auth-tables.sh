#! /usr/bin/env sh
echo "creating tables schema ${ZK_AUTH_DB_NAME}"

PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\" WITH SCHEMA public;"
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "SET search_path TO ${ZK_AUTH_DB_NAME},public;"
# integers are 8 bytes 
# y values are stored as 256 bit integers - 2^256 is 78 digits
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "CREATE TABLE IF NOT EXISTS ${ZK_AUTH_DB_NAME}.auth
        (id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
        user_id TEXT NOT NULL UNIQUE,
        y1 numeric(78, 0) NOT NULL,
        y2 numeric(78, 0) NOT NULL,
        created_at timestamp without time zone DEFAULT current_timestamp,
        updated_at timestamp without time zone);"



# updated_at date trigger
read -r -d '' DATE_TRIGGER <<- EOM
create or replace function updated_at_column()
returns trigger as \$$
begin
    new.updated_at = now();
    return new;
end;
\$$ language 'plpgsql';
EOM
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "${DATE_TRIGGER}"
PGPASSWORD=${ZK_AUTH_DB_ADMIN_PASSWORD} psql -d ${ZK_AUTH_DB_NAME} -U${ZK_AUTH_DB_ADMIN} -c "create trigger updated_at_trigger before insert or update on ${ZK_AUTH_DB_NAME}.auth for each row execute procedure updated_at_column();"
