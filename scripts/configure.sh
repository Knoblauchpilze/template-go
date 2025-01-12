#!/bin/bash

if [ "$#" -ne 1 ]; then
  echo "usage: ./configure.sh name-of-the-new-project"
  exit 1
fi

PROJECT=$1

echo "Starting configuration for ${PROJECT}"

echo "Rewriting go.mod..."
sed -i "s/template-go/${PROJECT}/g" go.mod

echo "Rewriting docker config..."
sed -i "s/template-go/${PROJECT}/g" build/template-go/Dockerfile
mv build/template-go build/${PROJECT}

echo "Creating main service from template app..."
sed -i "s/template-go/${PROJECT}/g" cmd/app/main.go
sed -i "s/app/${PROJECT}/g" cmd/app/Makefile
mv cmd/app cmd/${PROJECT}

echo "Updating build system..."
sed -i "s/template-go/${PROJECT}/g" Makefile
sed -i "s/template-frontend/${PROJECT}/g" .github/workflows/build-and-push.yml

echo "Updating README"
sed -i "s/\${placeholder-for-name}/${PROJECT}/g" README.template.md
mv README.template.md README.md

echo "Initializing git repository..."
rm -rf .git
git init

echo "All set! Don't forget to update the README and run the configure-database.sh script!"

