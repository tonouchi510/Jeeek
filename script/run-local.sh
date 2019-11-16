#!/usr/bin/env bash

export FIREBASE_CREDENTIALS="firebase-service-key-dev.json"
export FIREBASE_APIKEY="AIzaSyAF9quvSMu9n3LMWBrXw_aO5LYwBzqT4Gw"
export GOOGLE_APPLICATION_CREDENTIALS="gae-service-key-dev.json"
dev_appserver.py app.local.yaml
