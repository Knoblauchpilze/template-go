#!/bin/bash

if [ "$#" -ne 1 ]; then
  echo "usage: ./configure-database.sh name-of-the-new-database"
  exit 1
fi

DATABASE=$1

# TODO: Handle this
