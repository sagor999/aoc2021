GOOS ?= linux
GOARCH ?= amd64

export GOFLAGS = -mod=vendor

.PHONY: build
build:
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(GOFLAGS) -ldflags "-s -w" \
		-o build/aoc2021-$(GOARCH)

.PHONY: run
run: build
	./build/aoc2021-$(GOARCH)