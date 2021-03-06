BINARY_NAME := daily-prog
PKG := github.com/wizardsoftheweb/daily-programmer-cli
SEMVER := $(shell git describe --abbrev=0)
VERSION := $(shell git describe --always --long --dirty)

clean:
	@ rm -rf ./build

build-version:
	go build -i -v -o "build/${BINARY_NAME}" -ldflags="-X ${PKG}/cmd.PackageVersion=${VERSION}" ${PKG}

build: clean build-version

init:
	go mod init github.com/wizardsoftheweb/daily-programmer-cli

test:
	go test -v ./... -cover -race -coverprofile=.coverage.out

coverage: test
	go tool cover -func=.coverage.out
