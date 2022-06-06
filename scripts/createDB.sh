#!/usr/bin/env bash

CreateDB() {
   echo "Creating DB $1"
   sudo -u postgres psql -c "CREATE ROLE $1admin WITH LOGIN PASSWORD 'b2Led2ke';"
   sudo -u postgres psql -c "CREATE DATABASE $1;"
   sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE $1 TO $1admin;"
   # echo "CREATE ROLE $1admin WITH LOGIN PASSWORD 'b2Led2ke';" >> container/init.sql
   # echo "CREATE DATABASE $1;" >> container/init.sql
   # echo "GRANT ALL PRIVILEGES ON DATABASE $1 TO $1admin;" >> container/init.sql
   # echo "" >> container/init.sql
}

sudo systemctl start postgresql
CreateDB application
CreateDB auth
CreateDB company
CreateDB rc
CreateDB student
