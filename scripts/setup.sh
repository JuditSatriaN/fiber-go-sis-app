#!/bin/sh

dbname="db_pos_app"
port=${PGPORT:-7000}
user="postgres"
password="root"
host=${PGHOST:-localhost}

PGPASSWORD=$password psql -U $user -d "postgres" -h $host -p $port -c "CREATE DATABASE $dbname" 2> /dev/null
for filename in schema/postgres/*.sql; do
    PGPASSWORD=$password psql -h $host -p $port -d $dbname -U $user -f "$filename"
done