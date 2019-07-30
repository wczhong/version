# These are the values we want to pass for Version and BuildTime
GIT_TAG=`git describe --tags --always`
BUILD_TIME=`date +%FT%T%z`

VER_PACKAGE=example/version/version
# Setup the -ldflags option for go build here, interpolate the variable values 
# path must be project var path
LDFLAGS=-ldflags "-X ${VER_PACKAGE}.GitTag=${GIT_TAG} -X ${VER_PACKAGE}.BuildTime=${BUILD_TIME}"

server:
	go build ${LDFLAGS} -o bin/server main.go