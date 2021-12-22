#!/bin/bash
#
PORT=$1

if [ $PORT -eq ]; then
  PORT="80"
fi

echo "# BUSINESS - READ ALL #"
curl -i -X GET "localhost:$PORT/business"
echo
echo ------------
