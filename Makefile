BINARY_NAME=hardware-monitor
VERSION=$(shell git describe --tags --always --dirty)
BUILD_TIME=$(shell date +%FT%T%z)
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

.PHONY: all build clean test run install

all: build

build:
	go build ${LDFLAGS} -o ${BINARY_NAME} .

clean:
	go clean
	rm -f ${BINARY_NAME}
	rm -f ${BINARY_NAME}-*

test:
	go test -v ./...

run: build
	./${BINARY_NAME}

install:
	go install ${LDFLAGS} .

# Cross compilation
build-linux:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY_NAME}-linux-amd64 .


build-darwin:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY_NAME}-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${BINARY_NAME}-darwin-arm64 .

build-windows:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY_NAME}-windows-amd64.exe .

build-all: build-linux build-darwin build-windows
