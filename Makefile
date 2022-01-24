GIT_VERSION ?= $(shell git describe --tags --always --dirty)

SOURCE_DATE_EPOCH ?= $(shell git log -1 --pretty=%ct)
DATE_FMT = +'%Y-%m-%dT%H:%M:%SZ'
ifdef SOURCE_DATE_EPOCH
    BUILD_DATE ?= $(shell date -u -d "@$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u -r "$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u "$(DATE_FMT)")
else
    BUILD_DATE ?= $(shell date "$(DATE_FMT)")
endif

PKG=github.com/racing-telemetry/f1-dump/cmd

LDFLAGS="-X $(PKG).GitCommitSHA=$(GIT_VERSION) -X $(PKG).BuildDate=$(BUILD_DATE) -s -w"

build:
	CGO_ENABLED=0 go build -ldflags $(LDFLAGS) -o f1-dump
