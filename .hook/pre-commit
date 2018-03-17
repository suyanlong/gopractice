#!/usr/bin/env bash
set -e
set -o pipefail

go version
if [ $? != 0 ]; then
	printf "go not available?\n"
	exit 1
fi

if [ $(gofmt -l ./.. | wc -l) != 0 ]; then
    gofmt -d ./..
    exit 1
fi
