# PACKAGES=$(shell find . -name '*.go' -print0 | xargs -0 -n1 dirname | sort --unique)
export GO111MODULE=on
export GOPROXY=https://mod.yunzhanghu.com

GIT_TAG=1 #`git describe --tags --always`
BUILD_TIME=2 #`date +%FT%T%z`

VER_PACKAGE=/Users/yunzhanghu396/go/src/example/mysql
VER_FLAGS=-X ${VER_PACKAGE}.GitTag=${GIT_TAG} -X ${VER_PACKAGE}.BuildTime=${BUILD_TIME}

LDFLAGS=-ldflags "-w -s ${VER_FLAGS}"

tidy:
	go mod tidy -v

dep:
	go mod download

server:
	go build ${LDFLAGS} -o bin/server main.go  