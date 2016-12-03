OUT := cobra-viper-sample
PKG := github.com/sgykfjsm/cobra-viper-sample
VERSION := $(shell git describe --always --long)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)
GLIDE := $(shell which glide || :)

all: build

glide:
	@if [ -z "${GLIDE}" ]; then \
		go get -u -v github.com/Masterminds/glide ; \
	fi

deps: glide
	@glide install

test:
	@go test -v ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

static: vet lint

build: deps static
	go build -i -v -o ${OUT} -tags netgo -ldflags="-extldflags \"-static\" -w -s -X main.version=${VERSION}" ${PKG}

clean:
	@go clean -x

.PHONY: run server static vet lint
