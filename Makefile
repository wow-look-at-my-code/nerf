.PHONY: all test current-triplet
.EXTRA_PREREQS := Makefile

# Detect current platform
CURRENT_OS := $(shell go env GOOS)
CURRENT_ARCH := $(shell go env GOARCH)

# All Go source files
GO_FILES := $(shell find src -name '*.go')
export GO_FILES

all: current-triplet

current-triplet: bin/$(CURRENT_OS)-$(CURRENT_ARCH)/nerf

bin/linux-amd64/nerf:
	$(MAKE) -f src/build/platform.mk GOOS=linux GOARCH=amd64

bin/darwin-arm64/nerf:
	$(MAKE) -f src/build/platform.mk GOOS=darwin GOARCH=arm64
	ln -sfn darwin-arm64 bin/macos-arm64

bin/windows-amd64/nerf:
	$(MAKE) -f src/build/platform.mk GOOS=windows GOARCH=amd64

test: bin/linux-amd64/nerf
	chronic docker compose -f test/docker-compose.yml build
	docker compose -f test/docker-compose.yml run --rm test
