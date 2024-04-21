VERSION := $(shell git describe --tags --always)
GITREV := $(shell git rev-parse --short HEAD)
GITBRANCH := $(shell git rev-parse --abbrev-ref HEAD)
DATE := $(shell LANG=US date +"%a, %d %b %Y %X %z")

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/dist
GOARCH := $(ARCH)
GOENVVARS := GOBIN=$(GOBIN) CGO_ENABLED=1 GOOS=$(OS) GOARCH=$(GOARCH)
GOBINARY := lagrange-node
GOCMD := $(GOBASE)/cmd/baseapp/
SCRIPTS_FOLDER=$(GOBASE)/scripts

LDFLAGS += -X 'github.com/Lagrange-Labs/lagrange-node.Version=$(VERSION)'
LDFLAGS += -X 'github.com/Lagrange-Labs/lagrange-node.GitRev=$(GITREV)'
LDFLAGS += -X 'github.com/Lagrange-Labs/lagrange-node.GitBranch=$(GITBRANCH)'
LDFLAGS += -X 'github.com/Lagrange-Labs/lagrange-node.BuildDate=$(DATE)'

STOP := docker compose down --remove-orphans

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


# Linting, Teseting, Benchmarking
golangci_lint_cmd=github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2

install-linter:
	@echo "--> Installing linter"
	@go install $(golangci_lint_cmd)

lint:
	@echo "--> Running linter"
	@ $$(go env GOPATH)/bin/golangci-lint run --timeout=10m
.PHONY:	lint install-linter

test: stop
	docker compose -f docker-compose.yml up -d mongo
	docker compose -f docker-compose.yml up -d lagrangesc
	sleep 3
	docker compose -f docker-compose.yml up -d simavs-sync
	sleep 2
	docker ps -a
	trap '$(STOP)' EXIT; go test ./... --timeout=10m
.PHONY: test

run-db-mongo:
	docker compose -f docker-compose.yml up -d mongo
.PHONY: run-db-mongo

run-lagrange-sc:
	docker compose -f docker-compose.yml up -d lagrangesc
.PHONY: run-lagrange-sc

run-avs-sync:
	docker compose -f docker-compose.yml up -d simavs-sync
.PHONY: run-avs-sync

benchmark: 
	go test -run=NOTEST -timeout=30m -benchmem  -bench=. ./...
.PHONY: benchmark

# Local testnet
create-keystore:
	go run ./testutil/chainconfig/cmd/...

localnet-start: stop 
	docker compose -f docker-compose.yml up -d mongo
	docker compose -f docker-compose.yml up -d lagrangesc
	sleep 3
	docker compose -f docker-compose.yml up -d simserver
	docker compose -f docker-compose.yml up -d simsequencer
	docker compose -f docker-compose.yml up -d simavs-sync
	sleep 3
	docker compose -f docker-compose.yml up -d simnode1
	docker compose -f docker-compose.yml up -d simnode2
	docker compose -f docker-compose.yml up -d simnode3
	docker compose -f docker-compose.yml up -d simnode4
	docker compose -f docker-compose.yml up -d simnode5
	docker compose -f docker-compose.yml up -d simnode6
	docker compose -f docker-compose.yml up -d simnode7

stop:
	$(STOP)

.PHONY: create-keystore localnet-start stop

# Useful and Test Scripts
scgen: # Generate the go bindings for the smart contracts
	@ cd scinterface && sh generator.sh

register-operator: # Register an operator
	go run ./testutil/registerops/

.PHONY: scgen register-operator

# Run Components
run-server:
	go run ./cmd/baseapp/... run-server

run-client:
	go run ./cmd/baseapp/... run-client

run-sequencer:
	go run ./cmd/baseapp/... run-sequencer

.PHONY: run-server run-client run-sequencer