#!/usr/bin/env bash

echo Show parsed swagger difinition

cat static/swagger-ui/swagger/openapi.json | python -m json.tool
