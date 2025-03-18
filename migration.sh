#!/bin/bash
source docker.env

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "host=${PG_HOST} port=5432 dbname=${PG_DATABASE_NAME} user=${PG_USER} password=${PG_PASSWORD} sslmode=disable" up -v