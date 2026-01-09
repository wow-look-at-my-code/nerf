# Build rules - invoked with GOOS, GOARCH
SRCDIR := src
OUTDIR := bin/$(GOOS)-$(GOARCH)

COMMANDS := $(basename $(notdir $(wildcard $(SRCDIR)/commands/*.go)))
EXTRA_ALIASES := fd

all: $(OUTDIR)/nerf $(addprefix $(OUTDIR)/commands/,$(COMMANDS)) $(addprefix $(OUTDIR)/commands/,$(EXTRA_ALIASES))

$(OUTDIR)/nerf: $(GO_FILES)
	@mkdir -p $(dir $@)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $@ ./$(SRCDIR)/

$(addprefix $(OUTDIR)/commands/,$(COMMANDS) $(EXTRA_ALIASES)): $(OUTDIR)/nerf
	@mkdir -p $(dir $@)
	ln -sf ../nerf $@
