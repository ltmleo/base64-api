#!/bin/bash

set -e

API_URL="http://$1"

echo "Wating 404 in $API_URL"

for i in $(seq 10); do
    echo "Running test $i"
    status_code=$(curl -o /dev/null -s -w "%{http_code}\n" $API_URL)
    if [ $status_code -eq 404 ]; then
        echo "Test $i Succeeded"
        exit 0
    fi
    sleep 20
done