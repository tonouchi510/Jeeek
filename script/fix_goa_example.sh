#!/usr/bin/env bash

find . -maxdepth 1 -name "*.go" | while read fname; do
    if [[ ${fname} != "./main.go" && ${fname} != "./http.go" ]]; then
        mv -n ${fname} controller/;
    fi
done

find controller/*.go | xargs sed -i '' 's|package jeeek|package controller|g'