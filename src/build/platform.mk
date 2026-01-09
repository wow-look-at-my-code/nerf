# Build rules - invoked with GOOS, GOARCH
SRCDIR := src
OUTDIR := bin/$(GOOS)-$(GOARCH)

COMMANDS := $(basename $(notdir $(wildcard $(SRCDIR)/commands/*.go)))

all: $(OUTDIR)/nerf $(addprefix $(OUTDIR)/commands/,$(COMMANDS))

$(OUTDIR)/nerf: $(GO_FILES)
	@mkdir -p $(dir $@)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $@ ./$(SRCDIR)/

$(addprefix $(OUTDIR)/commands/,$(COMMANDS)): $(OUTDIR)/nerf
	@mkdir -p $(dir $@)
	ln -sf ../nerf $@
