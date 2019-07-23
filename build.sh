#!/bin/sh
# Sat Jul 16 22:15:04 CST 2019
#
version="2.1.$(git rev-list HEAD --count)-$(date +'%m%d%H')"
if [ "X$1" = "Xpkg" ]
then
  which zip || { echo 'zip command not find'; exit; }
  echo "Building Linux amd64 grf-$version"
  GOOS=linux GOARCH=amd64 go build -ldflags " -X main.version=$version"
  zip -m grf-$version-linux-amd64.zip grf
  
  echo "Building Macos amd64 grf-$version"
  GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$version"
  zip -m grf-$version-macos-amd64.zip grf
  
  echo "Building Windows amd64 grf-$version"
  GOOS=windows GOARCH=amd64 go build -ldflags " -X main.version=$version"
  zip -m grf-$version-win-x64.zip grf.exe
else
  echo "Building grf-$version"
  go build -ldflags "-X main.version=$version"
fi
