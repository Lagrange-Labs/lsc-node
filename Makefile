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

LDFLAGS += -X 'github.com/Lagrange-Labs/Lagrange-Node.Version=$(VERSION)'
LDFLAGS += -X 'github.com/Lagrange-Labs/Lagrange-Node.GitRev=$(GITREV)'
LDFLAGS += -X 'github.com/Lagrange-Labs/Lagrange-Node.GitBranch=$(GITBRANCH)'
LDFLAGS += -X 'github.com/Lagrange-Labs/Lagrange-Node.BuildDate=$(DATE)'


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

test:
	go test ./... --timeout=10m
.PHONY: test

# Local testnet
localnet-start: localnet-stop docker-build
	docker-compose up -d

localnet-stop:
	docker-compose down

.PHONY: localnet-build-nodes localnet-stop