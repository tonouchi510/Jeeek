#!/usr/bin/env bash

sed -i -e 's/localhost:8080/jeeek-dev.appspot.com/' static/swagger-ui/swagger/openapi.json
sed -i -e "s/\"schemes\":\[\"http\"\]/\"schemes\":\[\"https\"\]/g" static/swagger-ui/swagger/openapi.json
