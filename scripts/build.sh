#!/bin/bash --posix

ENV=$(uname -snr)
VER=$(cat .version)
UPT=$(date +"%Y%m%d")
TAG=$(echo $(git rev-parse --short HEAD)$([ -n "$(git status -s)"  ] && echo "-dev" || echo ""))

PROJECT=$1

SOURCE="cmd/${PROJECT}/main.go"
BINARY="bin/${PROJECT}"

echo "Making ${PROJECT}"
echo "Version:${TAG}"

go build -tags netgo                              \
    -installsuffix 'static'                       \
    -ldflags "                                    \
    -s -w                                         \
    -X '$(go list -m)/pkg/info.verStr=${VER}'     \
    -X '$(go list -m)/pkg/info.tagStr=${TAG}'     \
    -X '$(go list -m)/pkg/info.uptStr=${UPT}'     \
    -X '$(go list -m)/pkg/info.envStr=${ENV}'     \
    -X '$(go list -m)/pkg/info.AppName=${PROJECT}'\
    "                                             \
    -o ${BINARY} ${SOURCE}