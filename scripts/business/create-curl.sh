#!/bin/bash
#
PORT=$1

if [ $PORT -eq ]; then
  PORT="80"
fi

echo "# BUSINESS - CREATE #"
curl \
  -H 'Content-Type: application/json' \
  -i \
  -X POST \
  -d '{
  "name": "BusinessX",
  "phone": "+55 (11) 99999-9999",
  "address": {
    "zip_code": "01001-000"
  }
}' \
  "localhost:$PORT/business"
echo
echo ------------
