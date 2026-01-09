.PHONY: all test current-triplet linux-amd64 darwin-arm64 windows-amd64
.EXTRA_PREREQS := Makefile
.SECONDEXPANSION:

# Detect current platform
CURRENT_OS := $(shell go env GOOS)
CURRENT_ARCH := $(shell go env GOARCH)

# All Go source files
GO_FILES := $(shell find src -name '*.go')

# Commands derived from src/commands/*.go
SRCDIR := src
COMMANDS := $(basename $(notdir $(wildcard $(SRCDIR)/commands/*.go)))
# Only test go-derived commands; aliases are tested via their parent test files
BATS_FILES := $(addprefix $(SRCDIR)/commands/,$(addsuffix .bats,$(COMMANDS)))

# Function to get all symlinks for a platform: $(call SYMLINKS,triplet)
SYMLINKS = $(addprefix bin/$(1)/commands/,$(COMMANDS))

all: current-triplet

current-triplet: $(CURRENT_OS)-$(CURRENT_ARCH)

# Platform rules template - $(1)=GOOS, $(2)=GOARCH
define PLATFORM_RULES
bin/$(1)-$(2)/nerf: $$(GO_FILES)
	@mkdir -p $$(dir $$@)
	GOOS=$(1) GOARCH=$(2) go build -o $$@ ./$$(SRCDIR)/

bin/$(1)-$(2)/commands/%: bin/$(1)-$(2)/nerf
	@mkdir -p $$(dir $$@)
	ln -sf ../nerf $$@

endef

$(eval $(call PLATFORM_RULES,linux,amd64))
$(eval $(call PLATFORM_RULES,darwin,arm64))
$(eval $(call PLATFORM_RULES,windows,amd64))

# Triplet phony targets
linux-amd64 darwin-arm64 windows-amd64: %: bin/%/nerf $$(call SYMLINKS,%)

# Alias for macOS
bin/macos-arm64: bin/darwin-arm64/nerf
	ln -sfn darwin-arm64 bin/macos-arm64

test: linux-amd64
	chronic timeout 60 docker compose -f test/docker-compose.yml build
	timeout 60 docker compose -f test/docker-compose.yml run --rm test $(BATS_FILES)
