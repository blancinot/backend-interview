PACKAGE    = interview
DATE      ?= $(shell date +%FT%T%z)
VERSION   ?= $(shell echo $(shell cat $(PWD)/.version)-$(shell git describe --tags --always))

GO         = go
GOROOT     ?= $(shell go env GOROOT)
GODOC      = godoc
GOFMT      = gofmt

GOLINT     = golangci-lint

V          = 0
Q          = $(if $(filter 1,$V),,@)
M          = $(shell printf "\033[0;35m▶\033[0m")

GO_PACKAGE        = github.com/gustvision/powder-interview
API               = api
INTEGRATION       = integration


.PHONY: all
all: api

.PHONY: api
api:  ## Build api binary
	$(info $(M) building executable api…) @
	$Q cd cmd/$(API) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(API)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(API)_$(VERSION) bin/$(PACKAGE)_$(API)

.PHONY: integration
integration:  ## Build integration binary
	$(info $(M) building executable integration…) @
	$Q cd cmd/$(INTEGRATION) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(INTEGRATION)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(INTEGRATION)_$(VERSION) bin/$(PACKAGE)_$(INTEGRATION)

.PHONY: migrate
migrate:  ## Migrate sql db
	$(info $(M) migrate…) @
	$Q docker exec backend-interview-db psql --host=0.0.0.0 --user=postgres postgres -f migrations/00.sql

.PHONY: populate
populate:  ## Populate sql db with test data
	$(info $(M) populate…) @
	$Q docker exec backend-interview-db psql --host=0.0.0.0 --user=postgres postgres -f migrations/populate.sql

.PHONY: clean
clean:  ## Clean all test data in sql db
	$(info $(M) clean…) @
	$Q docker exec backend-interview-db psql --host=0.0.0.0 --user=postgres postgres -f migrations/clean.sql


# Vendor
.PHONY: vendor
vendor:
	$(info $(M) running go mod vendor…) @
	$Q $(GO) mod vendor

# Tidy
.PHONY: tidy
tidy:
	$(info $(M) running go mod tidy…) @
	$Q $(GO) mod tidy

# Check
.PHONY: check
check: vendor test lint

# Lint
.PHONY: lint
lint:
	$(info $(M) running $(GOLINT)…)
	$Q $(GOLINT) run

# Test
.PHONY: test
test:
	$(info $(M) running go test…) @
	$Q $(GO) test -cover -race -v ./...

# Helpers
.PHONY: go-version
go-version: ## Print go version used in this makefile
	$Q echo $(GO)

.PHONY: fmt
fmt: ## Format code
	$(info $(M) running $(GOFMT)…) @
	$Q $(GOFMT) ./...

.PHONY: doc
doc: ## Generate project documentation
	$(info $(M) running $(GODOC)…) @
	$Q $(GODOC) ./...

.PHONY: version
version: ## Print current project version
	@echo $(VERSION)

.PHONY: help
help: ## Print this
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
