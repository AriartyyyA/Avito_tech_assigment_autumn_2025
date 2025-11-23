#!/bin/sh

set -e

echo "Waiting for PostgreSQL..."
until pg_isready -h postgres -U postgres; do sleep 1; done

echo "Applying migrations..."
for file in /migrations/*.sql; do
  if [ -f "$file" ]; then
    echo "Applying: $(basename $file)"
    SQL=$(awk '
      $0 ~ /-- \+goose Up/ { in_up = 1; next }
      $0 ~ /-- \+goose Down/ { in_up = 0; next }
      $0 ~ /-- \+goose StatementBegin/ { next }
      $0 ~ /-- \+goose StatementEnd/ { next }
      in_up == 1 { print }
    ' "$file")
    
    if [ -n "$SQL" ]; then
      echo "$SQL" | psql -h postgres -U postgres -d avito_db -v ON_ERROR_STOP=1
      echo "✓ Migration $(basename $file) applied successfully"
    else
      echo "⚠ No SQL found in $(basename $file)"
    fi
  fi
done

echo "All migrations applied successfully!"

