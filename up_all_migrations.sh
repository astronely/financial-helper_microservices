#!/bin/sh
set -eu

DB_DSN="host=localhost port=54321 dbname=fhelper user=fhelper-user password=admin sslmode=disable"

find */migrations -type f -name '*.sql' | \
while IFS= read -r filepath; do
  echo "$(basename "$filepath") $filepath"
done | \
sort | \
while IFS=' ' read -r _ fullpath; do
  parent_dir=$(dirname "$fullpath")
#  echo "$parent_dir"
  goose -dir "$parent_dir" postgres "$DB_DSN" up-by-one
done
