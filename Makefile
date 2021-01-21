
BUILD_DIR=./bin
APP_NAME=diva-cli

VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

# Builds the project
build:
	go build ${LDFLAGS} -o ${BUILD_DIR}/${APP_NAME}

# Installs our project: copies binaries
install:
	go install ${LDFLAGS}

# Cleans our project: deletes binaries
clean:
	if [ -f ${BUILD_DIR}/${APP_NAME} ] ; then rm ${BUILD_DIR}/${APP_NAME} ; fi

.PHONY: clean install
