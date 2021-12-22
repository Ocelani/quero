#!/bin/bash
#
PORT=$1
ID=$2

if [ $PORT -eq ]; then
  PORT="80"
fi

if [ $ID -eq ]; then
  ID="1"
fi

echo "# BUSINESS - READ ONE #"
echo "search id: $ID"
curl -i -X GET "localhost:$PORT/business/$ID"
echo
echo ------------
