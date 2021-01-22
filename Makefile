
BUILD_DIR=./bin
APP_NAME=diva-cli

VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

# Builds the project
build:
	go build ${LDFLAGS} -o ${BUILD_DIR}/linux/${APP_NAME}
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/mac/${APP_NAME}
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/win/${APP_NAME}.exe

# Installs our project: copies binaries
install:
	go install ${LDFLAGS}

# Cleans our project: deletes binaries
clean:
	rm -rf ${BUILD_DIR}/

.PHONY: clean install
