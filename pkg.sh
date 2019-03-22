#!/bin/sh
# 2019-03-22
#
Version="1.0.$(git rev-list --all --count)"


echo "Building $Version Linux amd64 ..."
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=${Version}" -o grf
zip -m grf-$Version-linux-amd64.zip grf

echo "Building $Version Macos amd64 ..."
GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=${Version}" -o grf
zip -m grf-$Version-macos-amd64.zip grf

echo "Building $Version Windows amd64 ..."
GOOS=windows GOARCH=amd64 go build -ldflags "--X main.version=${Version}" -o grf.exe
zip -m grf-$Version-win-x64.zip grf.exe
