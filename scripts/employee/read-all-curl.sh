#!/bin/bash
#
PORT=$1

if [ $PORT -eq ]; then
  PORT="80"
fi

echo "# EMPLOYEE - READ ALL #"
curl -i -X GET "localhost:$PORT/employee"
echo
echo ------------
