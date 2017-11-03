#!/bin/bash

while true; do
    find controllers -name '*.go' | APP_ENV=test entr -dr go test -v ./controllers
    [ "$?" = "130" ] && break
done
