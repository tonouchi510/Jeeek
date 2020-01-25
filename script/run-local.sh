#!/usr/bin/env bash

if [ "$1" = "" ]
then
    echo "no argument"
    exit 1
fi
echo "$1"

export ADMIN_PASSWORD="$1"
export FIREBASE_CREDENTIALS="firebase-service-key-dev.json"
export FIREBASE_APIKEY="AIzaSyAF9quvSMu9n3LMWBrXw_aO5LYwBzqT4Gw"
export GOOGLE_APPLICATION_CREDENTIALS="gae-service-key-dev.json"
dev_appserver.py app.local.yaml
