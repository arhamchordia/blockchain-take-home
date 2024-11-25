BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')
PACKAGES_UNIT=$(shell go list ./... )

# don't override user values
ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match 2>/dev/null)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

# Update the ldflags with the app, client & server names
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=blog \
	-X github.com/cosmos/cosmos-sdk/version.AppName=blogd \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

###########
# Install #
###########

all: install

install:
	@echo "--> ensure dependencies have not been modified"
	@go mod verify
	@echo "--> installing blogd"
	@go install $(BUILD_FLAGS) -mod=readonly ./cmd/blogd

init:
	./scripts/init.sh

###########
#  Proto  #
###########

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh

###########
#   Test  #
###########

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock norace' $(PACKAGES_UNIT)


script-test: install
	@echo "Test blog"
	@$(protoImage) sh ./scripts/test_blog.sh

###########
#  Lint   #
###########

lint-format:
	@go run golangci-lint run ./... --fix