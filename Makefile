######################################################
# Docker Targets
# Needs docker installed on the system
######################################################

# Define variables for PostgreSQL container
POSTGRES_CONTAINER := postgres
POSTGRES_DB := cc_terrarium
POSTGRES_USER := postgres
FARM_DB_DUMP_FILE := $(POSTGRES_DB).psql
COVERAGE_FILE = coverage/coverage.txt

# Define variables for pg_dump command
DUMP_DIR := ./data

.PHONY: docker-init
docker-init:  ## Initialize the environment before running docker commands
	@touch ${HOME}/.netrc

.PHONY: db-dump
db-dump: start-db  ## Target for dumping PostgreSQL database to a file
	docker compose exec -T $(POSTGRES_CONTAINER) pg_dump --column-inserts -U $(POSTGRES_USER) -f /docker-entrypoint-initdb.d/$(FARM_DB_DUMP_FILE) -Fc $(POSTGRES_DB)

.PHONY: db-update  ## Restore database from the database bump file. Ignore errors for existing rows.
db-update: data/$(FARM_DB_DUMP_FILE) start-db
	docker compose exec -T $(POSTGRES_CONTAINER) pg_restore -a -d $(POSTGRES_DB) -U $(POSTGRES_USER) /docker-entrypoint-initdb.d/$(FARM_DB_DUMP_FILE) || echo ignore errors for already existing rows.

.PHONY: docker-build
docker-build:  ## Build API container image
	docker compose build

.PHONY: docker-run
docker-run:  ## Starts app in docker containers
	docker compose up -d

.PHONY: start-db
start-db:  ## Starts database in docker containers
	docker compose up -d postgres

.PHONY: docker-stop
docker-stop:  ## Stops and removes docker containers
	docker compose down

.PHONY: docker-stop-clean
docker-stop-clean:  ## Stops and removes containers as well as volumes to cleanup database
	docker compose down -v

docker-tools-build:  ## Build CLI & farm-harvester images
	docker compose --profile tooling build

docker-harvest: docker-tools-build start-db  ## Run harvest scripts on the Farm folder in a containerized environment
	docker compose run --rm farm-harvester

docker-seed: docker-harvest  ## DEPRECATED. alias to docker-harvest target
	@echo "\nNOTE: This target is deprecated. Use make docker-harvest instead"

docker-api-test: ## Run API unit tests in a containerized environment
	docker compose run --build --rm terrarium-unit-test

proto:
	make proto -f scripts/protoc.mak

######################################################
# Go Targets
# Needs Go installed on the system
######################################################

-include .env
export

CLI_SRCS := $(shell find ./src/pkg ./src/cli \( -name '*.go' -o -name 'go.mod' \))
BUILD_DATE=$(shell date "+%Y-%m-%d")
BUILD_DIR=.bin
CLI_NAME = terrarium
BINARY_NAME = $(BUILD_DIR)/$(CLI_NAME)
BINARY_NAME_WIN = $(BUILD_DIR)/win/$(CLI_NAME).exe
BINARY_NAME_LINUX = $(BUILD_DIR)/linux/$(CLI_NAME)
BINARY_NAME_MACOS_ARM = $(BUILD_DIR)/macos/arm64/$(CLI_NAME)
BINARY_NAME_MACOS_I386 = $(BUILD_DIR)/macos/i386/$(CLI_NAME)

ifeq (${TAG},)
	TAG=$(shell git describe --exact-match --tags 2> /dev/null)
	VERSION=$(shell git rev-parse --short HEAD 2> /dev/null)
	ZIP_DATE=$(shell date "+%Y%m%d")
	ZIP_FILE=$(BUILD_DIR)/terrarium_$(ZIP_DATE)_$(VERSION).zip
else
	VERSION=$(TAG)
	ZIP_FILE=$(BUILD_DIR)/terrarium_$(TAG).zip
	VERFILE=$(shell echo $(TAG) > $(BUILD_DIR)/latest_version.txt)
endif

GO_LDFLAGS := -s -w
GO_LDFLAGS := -X github.com/cldcvr/terrarium/src/cli/internal/build.Version=$(VERSION) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/cldcvr/terrarium/src/cli/internal/build.Date=$(BUILD_DATE) $(GO_LDFLAGS)
GO_LDFLAGS := -ldflags "$(GO_LDFLAGS)"

.PHONY: mod-clean
mod-clean:  ## delete go*.sum files
	@echo "deleting .sum files..."
	@rm -f ./src/api/go.sum ./src/cli/go.sum ./src/pkg/go.sum ./go.work.sum

.PHONY: clean-binaries
clean-binaries:  ## Clean up artifacts generated by make
	@echo "-- Cleaning build directory"
	@rm -rf $(BUILD_DIR)/

.PHONY: mod-tidy
mod-tidy:  ## run go mod tidy on each workspace entity, and then sync workspace
	@echo "running tidy on api go module..."
	@cd src/api && go mod tidy -e
	@echo "running tidy on cli go module..."
	@cd src/cli && go mod tidy -e
	@echo "running tidy on pkg go module..."
	@cd src/pkg && go mod tidy -e
	@echo "running sync on go workspace..."
	@go mod download && go work sync

.PHONY: test
test:  ## Run go unit tests
	mkdir -p coverage
	go test -v -tags=mock -coverprofile $(COVERAGE_FILE) github.com/cldcvr/terrarium/...
	@echo "-- Test coverage for terrarium --"
	@go tool cover -func=$(COVERAGE_FILE)|grep "total:"

define make_binary
	@echo "-- Building application binary $(1) for $(2)-$(3)"
	CGO_ENABLED=0 GOOS=$(2) GOARCH=$(3) go build $(GO_LDFLAGS) -v -o $(1) ./src/cli/terrarium
endef

$(BINARY_NAME): $(CLI_SRCS)
	$(call make_binary, $(BINARY_NAME),$(shell go env GOOS),$(shell go env GOARCH))
.PHONY: binary
binary: $(BINARY_NAME)  ## Build application binary (native)

$(BINARY_NAME_WIN): $(CLI_SRCS)
	$(call make_binary,$@,windows,amd64)
.PHONY: binary_win
binary_win: $(BINARY_NAME_WIN)  ## Build application binary for Windows

$(BINARY_NAME_LINUX): $(CLI_SRCS)
	$(call make_binary,$@,linux,amd64)
.PHONY: binary_linux
binary_linux: $(BINARY_NAME_LINUX)  ## Build application binary for Linux

$(BINARY_NAME_MACOS_ARM): $(CLI_SRCS)
	$(call make_binary,$@,darwin,arm64)
$(BINARY_NAME_MACOS_I386): $(CLI_SRCS)
	$(call make_binary,$@,darwin,amd64)
.PHONY: binary_macos
binary_macos: $(BINARY_NAME_MACOS_ARM) $(BINARY_NAME_MACOS_I386)  ## Build application binaries for MacOS

$(ZIP_FILE): $(BINARY_NAME_WIN) $(BINARY_NAME_LINUX) binary_macos
	cd $(BUILD_DIR); \
	zip -r ../$(ZIP_FILE) *
.PHONY: binaries
binaries: $(ZIP_FILE)  ## Build binary for each supported platform and archive to zip

.PHONY: install
install:  ## Install the CLI native binary into GOBIN
	@echo "-- Installing native binary in $(shell go env GOBIN)"
	go install $(GO_LDFLAGS) github.com/cldcvr/terrarium/src/cli/terrarium

######################################################
# Farm Targets
# Needs terraform installed on the system
######################################################

ifeq ($(FARM_DIR),)
FARM_DIR := examples/farm
endif

.PHONY: farm-harvest
farm-harvest: farm-resource-harvest farm-module-harvest farm-mapping-harvest  ## Run all harvest commands on the farm directory

.PHONY: farm-resource-harvest  ## Harvest terraform provider resources from module list file
farm-resource-harvest: $(FARM_DIR)/modules.yaml
	terrarium harvest resources --module-list-file $(FARM_DIR)/modules.yaml

.PHONY: farm-module-harvest  ## Harvest terraform modules from module list file
farm-module-harvest: $(FARM_DIR)/modules.yaml
	terrarium harvest modules --module-list-file $(FARM_DIR)/modules.yaml

.PHONY: farm-mapping-harvest  ## Harvest attribute mappings from module list file
farm-mapping-harvest: $(FARM_DIR)/modules.yaml
	terrarium harvest mappings --module-list-file $(FARM_DIR)/modules.yaml

######################################################
# Farm releases pull
# Needs access to the farm repo
######################################################

FARM_REPO := github.com/cldcvr/terrarium-farm

ifeq ($(FARM_VERSION),)
FARM_VERSION := latest
endif

data/$(FARM_DB_DUMP_FILE):
	$(call farm-release-pull)

.PHONY: farm-release-pull
farm-release-pull:
	@echo "Downloading db dump from the latest terrarium-farm release..."
	@gh release download $(FARM_VERSION) --clobber -p 'terrarium_farm.psql' -O data/$(FARM_DB_DUMP_FILE) -R $(FARM_REPO)

.PHONY: help
help:
	@grep -hE '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
