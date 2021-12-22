#!/bin/bash
#
PORT=$1

if [ $PORT -eq ]; then
  PORT="80"
fi

echo "# EMPLOYEE - CREATE #"
curl \
  -H 'Content-Type: application/json' \
  -i \
  -X POST \
  -d '{
  "name": "Employee",
  "sallary": "R$ 1.000,00",
  "role": "Developer",
}' \
  "localhost:$PORT/employee"
echo
echo ------------
