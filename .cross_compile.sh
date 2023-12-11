#!/usr/bin/env bash

set -e

DIST_PREFIX="go-common"

TARGET_DIR="dist"

if [ -n "${1}" ]
then
PLATFORMS="$1"
else
PLATFORMS="darwin/amd64 darwin/arm64  linux/386 linux/amd64 linux/arm linux/arm64 linux/mips linux/mips64 linux/riscv64 windows/amd64 windows/386 windows/arm windows/arm64"
fi
echo "$PLATFORMS"

BUILD_VERSION=$(cat version)
BUILD_DATE=$(date "+%F %T")
COMMIT_SHA1=$(git rev-parse HEAD)

rm -rf ${TARGET_DIR}
mkdir ${TARGET_DIR}

for pl in ${PLATFORMS}; do
    export GOOS=$(echo ${pl} | cut -d'/' -f1)
    export GOARCH=$(echo ${pl} | cut -d'/' -f2)
    export TARGET=${TARGET_DIR}/${DIST_PREFIX}_${GOOS}_${GOARCH}
    if [ "${GOOS}" == "windows" ]; then
        export TARGET=${TARGET_DIR}/${DIST_PREFIX}_${GOOS}_${GOARCH}.exe
    fi
    if [ "${GOOS}" == "linux" ]; then
        export CGO_CFLAGS="-g -O2 -Wno-return-local-addr"  
    fi

    echo "build => ${TARGET}"
    go build -trimpath -o ${TARGET} -ldflags "-X 'github.com/clibing/go-common/cmd.buildVersion=${BUILD_VERSION}' -X 'github.com/clibing/go-common/cmd.buildDate=${BUILD_DATE}' -X 'github.com/clibing/go-common/cmd.commitId=${COMMIT_SHA1}' -X 'github.com/clibing/go-common/cmd.execBinName=${DIST_PREFIX}' -w -s"
done
