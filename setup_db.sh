#!/usr/bin/env bash

sqlite3 data.db <<EOF
.read setup.sql
EOF

jq -c '.[]' godowns.json | while read row; do
  sqlite3 data.db "INSERT INTO godowns (id, name, parent_godown) VALUES (
    '$(echo $row | jq -r '.id')',
    '$(echo $row | jq -r '.name')',
    '$(echo $row | jq -r '.parent_godown')'
  );"
done

jq -c '.[]' items.json | while read row; do
  sqlite3 data.db "INSERT INTO items (item_id, name, quantity, category, price, status, godown_id, brand, attributes, image_url) VALUES (
    '$(echo $row | jq -r '.item_id')',
    '$(echo $row | jq -r '.name')',
    $(echo $row | jq -r '.quantity'),
    '$(echo $row | jq -r '.category')',
    $(echo $row | jq -r '.price'),
    '$(echo $row | jq -r '.status')',
    '$(echo $row | jq -r '.godown_id')',
    '$(echo $row | jq -r '.brand')',
    '$(echo $row | jq -c '.attributes')',
    '$(echo $row | jq -r '.image_url')'
  );"
done