#!/bin/bash

# PUT YOUR API KEY HERE
api_key=YOUR_API_KEY


if [ $# -ne 1 ]; then
    echo "usage: $0 <email>"
    exit 1
fi

email=$1

res_json=$(curl "https://api.hubuco.com/api/v3/?api=$api_key&email=$email" 2>/dev/null)

resultcode=$(echo $res_json | sed -r 's/.*resultcode[\": ]+([0-9]*).*/\1/')
error=$(echo $res_json | sed -r 's/.*error[\": ]+([^\"\}]*).*/\1/')

case "$resultcode" in
    1)  echo Ok ;;
    2)  echo Catch All ;;
    3)  echo Unknown ;;
    4)  echo Error: $error ;;
    5)  echo Disposable ;;
    6)  echo Invalid ;;
esac

