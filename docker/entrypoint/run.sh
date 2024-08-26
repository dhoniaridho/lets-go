#!/bin/sh
printenv > .env
goose -dir ./db/migrations postgres $DB_URL up
./main.exe
