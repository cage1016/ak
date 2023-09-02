GO ?= go
EXECUTABLE := ak
GOFILES := $(shell find . -type f -name "*.go")
TAGS ?=
LDFLAGS ?= -X 'github.com/cage1016/ak/cmd.Version=$(VERSION)' -X 'github.com/cage1016/ak/cmd.Commit=$(COMMIT)'

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

ifneq ($(DRONE_TAG),)
	VERSION ?= $(DRONE_TAG)
else
	VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)
endif
COMMIT ?= $(shell git rev-parse --short HEAD)

# Regenerates OPA data from rego files
HAVE_GO_BINDATA := $(shell command -v go-bindata 2> /dev/null)
generate: ## go generate
ifndef HAVE_GO_BINDATA
	@echo "requires 'go-bindata' (go get -u github.com/kevinburke/go-bindata/go-bindata)"
	@exit 1 # fail	
else
	go generate ./...
endif

.PHONY: release
release: ## release
	goreleaser release --skip-publish --rm-dist	

.PHONY: run-image
run-image: ## pack build
	docker build -t ghcr.io/cage1016/ak-run-image:0.1.0 -f run.Dockerfile .
	docker push ghcr.io/cage1016/ak-run-image:0.1.0

.PHONY: build
build: generate ## Build container image
	pack build ghcr.io/cage1016/ak:0.1.0 --builder gcr.io/buildpacks/builder:v1 --run-image ghcr.io/cage1016/ak-run-image:0.1.0 --env GOOGLE_RUNTIME_VERSION=1.20
	# docker push ghcr.io/cage1016/ak:0.1.0

.PHONY: 
t: generate ## Run t
	rm -rf test
	mkdir test
	go run main.go --folder test init
	go run main.go --folder test new cmd
	go run main.go --folder test add ga -s
	go run main.go --folder test add l

.PHONY: help
help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_0-9-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help