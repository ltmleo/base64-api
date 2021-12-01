#!/bin/bash

set -e

if [ ! -n "$1" ]; then
    echo "Missing API URL"
    exit 1
fi

API_URL="http://$1/api/healthcheck"

echo "Wating 200 in $API_URL for 200s"

for i in $(seq 100); do
    echo "Running test $i"
    status_code=$(curl -o /dev/null -s -w "%{http_code}\n" $API_URL)
    if [ $status_code -eq 200 ]; then
        echo "Test $i Succeeded"
        exit 0
    fi
    sleep 2
done

exit 1