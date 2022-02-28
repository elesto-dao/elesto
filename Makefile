PACKAGES="./x/..."

# build paramters
BUILD_FOLDER = build
APP_VERSION = $(git describe --tags --always)

VERSION := $(shell echo $(shell git describe --always --match "v*") | sed 's/^v//')
TMVERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
BINDIR ?= $(GOPATH)/bin
BUILDDIR ?= $(CURDIR)/build
HTTPS_GIT := https://github.com/elesto-dao/elesto.git
DOCKER := $(shell which docker)

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += gcc
endif

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=elesto \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=elestod \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
			-X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TMVERSION)

ifeq ($(ENABLE_ROCKSDB),true)
  BUILD_TAGS += rocksdb_build
  test_tags += rocksdb_build
else
  $(warning RocksDB support is disabled; to build and test with RocksDB support, set ENABLE_ROCKSDB=true)
endif

# DB backend selection
ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (badgerdb,$(findstring badgerdb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb
  BUILD_TAGS += badgerdb
endif
# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(COSMOS_BUILD_OPTIONS)))
  ifneq ($(ENABLE_ROCKSDB),true)
    $(error Cannot use RocksDB backend unless ENABLE_ROCKSDB=true)
  endif
  CGO_ENABLED=1
  BUILD_TAGS += rocksdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif
# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(COSMOS_BUILD_OPTIONS)))
  BUILD_TAGS += boltdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb
endif

ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

# Check for debug option
ifeq (debug,$(findstring debug,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -gcflags "all=-N -l"
endif

all: build lint test

###############################################################################
###                                  Build                                  ###
###############################################################################

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/
build-linux:
	GOOS=linux GOARCH=amd64 LEDGER_ENABLED=false $(MAKE) build

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/


clean:
	@echo clean build folder $(BUILD_FOLDER)
	rm -rf $(BUILD_FOLDER)
	@echo done

.PHONY: build build-linux clean

###############################################################################
###                          Tools & Dependencies                           ###
###############################################################################

go.sum: go.mod
	echo "Ensure dependencies have not been modified ..." >&2
	go mod verify
	go mod tidy

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

test:
	@go test -mod=readonly $(PACKAGES) -cover -race


# TODO: multiple simulations can be run, this need to test more
sim:
	@go test -benchmem -bench BenchmarkSimulation ./app -NumBlocks=200 -BlockSize 50 -Commit=true -Verbose=true -Enabled=true

###############################################################################
###                                Protobuf                                 ###
###############################################################################

protogen:
	starport generate proto-go

###############################################################################
###                              Documentation                              ###
###############################################################################

docs:
	@echo "launch local documentation portal"
	mkdocs serve

openapi:
	starport generate openapi

.PHONY: docs openapi

###############################################################################
###                                Linting                                  ###
###############################################################################

lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

.PHONY: lint

###############################################################################
###                           Chain Initialization                          ###
###############################################################################

start-dev: install
	./scripts/seeds/00_start_chain.sh

seed:
	./scripts/seeds/01_identifier_seeds.sh

###############################################################################
###                                CI / CD                                  ###
###############################################################################

# TODO: running this with -race options causes problems in the cli tests
test-ci:
	go test -coverprofile=coverage.txt -covermode=atomic -mod=readonly $(PACKAGES)

###############################################################################
###                                RELEASE                                  ###
###############################################################################

changelog:
	git-chglog --output CHANGELOG.md

_get-release-version:
ifneq ($(shell git branch --show-current | head -c 9), release/v)
	$(error this is not a release branch. a release branch should be something like 'release/v1.2.3')
endif
	$(eval APP_VERSION = $(subst release/,,$(shell git branch --show-current)))
#	@echo -n "releasing version $(APP_VERSION), confirm? [y/N] " && read ans && [ $${ans:-N} == y ]

release-prepare: _get-release-version
	@echo making release $(APP_VERSION)
ifndef APP_VERSION
	$(error APP_VERSION is not set, please specifiy the version you want to tag)
endif
	git tag $(APP_VERSION)
	git-chglog --output CHANGELOG.md
	git tag $(APP_VERSION) --delete
	git add CHANGELOG.md && git commit -m "chore: update changelog for $(APP_VERSION)"
	@echo release complete

git-tag:
ifndef APP_VERSION
	$(error APP_VERSION is not set, please specifiy the version you want to tag)
endif
ifneq ($(shell git rev-parse --abbrev-ref HEAD),main)
	$(error you are not on the main branch. aborting)
endif
	git tag -s -a "$(APP_VERSION)" -m "Changelog: https://github.com/elesto-dao/elesto/blob/main/CHANGELOG.md"
