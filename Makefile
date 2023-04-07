VERSION := $(shell git describe --tags --always)
GITREV := $(shell git rev-parse --short HEAD)
GITBRANCH := $(shell git rev-parse --abbrev-ref HEAD)
DATE := $(shell LANG=US date +"%a, %d %b %Y %X %z")

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/dist
GOENVVARS := GOBIN=$(GOBIN) CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH)
GOBINARY := lagrange-node
GOCMD := $(GOBASE)/cmd/baseapp/
SCRIPTS_FOLDER=$(GOBASE)/scripts

LDFLAGS += -X 'github.com/Lagrange-Labs/lagrange-node.Version=$(VERSION)'
LDFLAGS += -X 'github.com/Lagrange-Labs/lagrange-node.GitRev=$(GITREV)'
LDFLAGS += -X 'github.com/Lagrange-Labs/lagrange-node.GitBranch=$(GITBRANCH)'
LDFLAGS += -X 'github.com/Lagrange-Labs/lagrange-node.BuildDate=$(DATE)'


# Building the docker image and the binary
build: ## Builds the binary locally into ./dist
	$(GOENVVARS) go build -ldflags "all=$(LDFLAGS)" -o $(GOBIN)/$(GOBINARY) $(GOCMD)
.PHONY: build

docker-build: ## Builds a docker image with the node binary
	docker build -t lagrange-node -f ./Dockerfile .
.PHONY: docker-build


# Protobuf
proto-gen:
	@ sh $(SCRIPTS_FOLDER)/proto-gen.sh
.PHONY: proto-gen


# Linting, Teseting
golangci_lint_cmd=github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

lint:
	@echo "--> Running linter"
	@ go run $(golangci_lint_cmd) run --timeout=10m
.PHONY:	lint

test: run-db-mongo
	go test ./... --timeout=10m
.PHONY: test

run-db-mongo: stop
	docker-compose -f docker-compose.yml up -d mongo
.PHONY: run-db-mongo

# Local testnet
localnet-start: stop docker-build
	docker-compose up -d

stop:
	docker-compose down --remove-orphans

.PHONY: localnet-start stop