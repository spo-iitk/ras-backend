#!/usr/bin/env bash

CreateDB() {
   echo "Creating DB $1"
   sudo -u postgres psql -c "CREATE ROLE $1admin WITH PASSWORD 'b2Led2ke';"
   sudo -u postgres psql -c "CREATE DATABASE $1;"
   sudo -u postgres psql -c "ALTER ROLE $1admin WITH LOGIN;"
   sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE $1 TO $1admin;"
}

sudo service postgresql start
CreateDB auth
CreateDB student
