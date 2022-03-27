
COMMIT_HASH := $(shell git rev-parse --short HEAD 2>/dev/null)
GIT_VERSION := $(shell git describe --tags --exact-match 2>/dev/null || git describe --tags 2>/dev/null || echo "v0.0.0-$(COMMIT_HASH)")
BUILD_DATE  := $(shell date +%FT%T%z)

dep:
	go mod tidy

PATTERN =
release: VERSION ?= $(shell echo $(GIT_VERSION) | sed 's/^v//' | awk -F'[ .]' '{print $(PATTERN)}')
release: PUSH    ?= false
release: ## Prepare release
	@ ./release.sh "$(VERSION)" "$(PUSH)" "$(GIT_VERSION)"


patch: PATTERN = '\$$1\".\"\$$2\".\"\$$3+1'
patch: release ## Prepare Patch release

minor: PATTERN = '\$$1\".\"\$$2+1\".0\"'
minor: release ## Prepare Minor release

major: PATTERN = '\$$1+1\".0.0\"'
major: release ## Prepare Major release


test:
	echo $(GIT_VERSION)
	echo $(COMMIT_HASH)
	echo $(BUILD_DATE)