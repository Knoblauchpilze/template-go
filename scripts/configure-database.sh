#!/bin/bash

if [ "$#" -ne 1 ]; then
  echo "usage: ./configure-database.sh name-of-the-new-database"
  exit 1
fi

DATABASE=$1

echo "Modifying database files..."
sed -i "s/template_service/${DATABASE}/g" database/template/db_user_create.sql
sed -i "s/template_service/${DATABASE}/g" database/template/db_user_drop.sql

sed -i "s/template_service/${DATABASE}/g" database/template/db_create.sql
sed -i "s/template_service/${DATABASE}/g" database/template/db_drop.sql

sed -i "s/template_service/${DATABASE}/g" database/template/migrations/1_create_initial_schema.up.sql

sed -i "s/template_service/${DATABASE}/g" database/Makefile
sed -i "s/template/${DATABASE}/g" database/Makefile

echo "Changing file structure..."
mv database/template "database/${DATABASE}"

echo "Updating CI workflows"
sed -i "s/template_service/${DATABASE}/g" .github/workflows/database-migration-tests.yml
sed -i "s/template/${DATABASE}/g" .github/workflows/database-migration-tests.yml

sed -i "s/template_service/${DATABASE}/g" .github/workflows/build-and-push.yml
sed -i "s/template\//${DATABASE}\//g" .github/workflows/build-and-push.yml
sed -i "s/template database/${DATABASE} database/g" .github/workflows/build-and-push.yml

echo "Updating test files..."
sed -i "s/template_service/${DATABASE}/g" internal/controller/helper_test.go

echo "Consolidating config files..."
sed -i "s/template_service/${DATABASE}/g" cmd/app/configs/config-prod.yml
sed -i "s/template_service/${DATABASE}/g" cmd/app/configs/config-template-dev.yml

sed -i "s/template_service/${DATABASE}/g" cmd/app/internal/config.go
sed -i "s/template_service/${DATABASE}/g" cmd/app/internal/config_test.go

echo "All set! Don't forget to update the README and run the configure.sh script!"
