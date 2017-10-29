#!/bin/bash

while true; do
    find . -name '*.go' -not -name '*_test.go' -not -ipath './goagen*' | entr -dr go run *.go
    [ "$?" = "130" ] && break
done
