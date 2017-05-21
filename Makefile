SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')
TESTABLE_PACKAGES := $(shell go list ./... | grep -v vendor)
TIMESTAMP := $(shell date +%s)

VERSION=0.1.0

LDFLAGS=-ldflags '-X main.version=${VERSION}'

BUILDDIR=build

BINARY=imagetoascii
DOCKER_BINARY=$(BINARY).docker

.DEFAULT_GOAL: $(BUILDDIR)/$(BINARY)

$(BUILDDIR)/$(BINARY): ${SOURCES}
	go build ${LDFLAGS} -o build/${BINARY} *.go

$(BUILDDIR)/$(BINARY).docker: ${SOURCES}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/${DOCKER_BINARY} *.go

docker: $(BUILDDIR)/$(DOCKER_BINARY)
	docker build -t ${BINARY}:${VERSION} -t ${BINARY} .

.PHONY: test
test:
	go test -v ${TESTABLE_PACKAGES}

.PHONY: clean
clean:
	@if [ -f "${BUILDDIR}/${BINARY}" ] ; then rm "${BUILDDIR}/${BINARY}" ; fi
