BUILD_VERSION   	:= $(shell cat version)
BUILD_DATE      	:= $(shell date "+%F %T")
COMMIT_SHA1     	:= $(shell git rev-parse HEAD)
DIST_PREFIX 		:= go-common

all: clean
	bash .cross_compile.sh

single: clean-single
	go build -trimpath -o "dist/single" -ldflags " -X 'github.com/clibing/go-common/cmd.buildVersion=${BUILD_VERSION}' -X 'github.com/clibing/go-common/cmd.buildDate=${BUILD_DATE}' -X 'github.com/clibing/go-common/cmd.commitId=${COMMIT_SHA1}' -X 'github.com/clibing/go-common/cmd.execBinName=${DIST_PREFIX}' -w -s"

release: all
	ghr -u clibing -t ${GITHUB_TOKEN} -replace -recreate -name "Bump ${BUILD_VERSION}" --debug ${BUILD_VERSION} dist

pre-release: all
	ghr -u clibing -t ${GITHUB_TOKEN} -replace -recreate -prerelease -name "Bump ${BUILD_VERSION}" --debug ${BUILD_VERSION} dist

install:
	go install -ldflags "-w -s -X 'github.com/clibing/go-common/cmd.buildVersion=${BUILD_VERSION}' -X 'github.com/clibing/go-common/cmd.buildDate=${BUILD_DATE}' -X 'github.com/clibing/go-common/cmd.commitId=${COMMIT_SHA1}' -X 'github.com/clibing/go-common/cmd.execBinName=${DIST_PREFIX}'" 
debug:
	go install -trimpath -gcflags "all=-N -l"

clean:
	rm -rf dist

clean-single:
	rm -rf dist/single

.PHONY: all release pre-release clean install debug
