TAG := $(shell git rev-parse --short HEAD --dirty)
.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST)  | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: build.darwin ## Build darwin binary , no testing
darwin: test build.darwin  ## Build a webserver debug binary on darwin
dev: test build.linux.debug	 ## Cross compile a linux binary from darwin in debug mode
debug: test build.linux.debug build.docker.debug ## Create a docker container with a debug mode compiled binary and source code. Expose the binary via gdbserver on port 1234. Tag and push docker to registry
release: test build.linux build.docker ## Cross compile a linux binary in release mode (Takes longer)
cross: cross.compile ## Add dependencies needed for cross compilation
run: run.debug ## Build and run a docker container locally with debugger on port 1234
test: ## Run cargo test on all integration tests
	@echo "Running tests"
	@cargo test

.PHONY: build.darwin
build.darwin: ## Build a debug binary on darwin
	@echo "Building binary for darwin"
	@cargo build

.PHONY: build.linux.debug
build.linux.debug:
	@echo "Build binary for linux kernel in debug mode"
	@TARGET_CC=x86_64-linux-musl-gcc cargo build --target=x86_64-unknown-linux-musl

.PHONY: build.linux
build.linux:
	@echo "Build binary for linux kernel in release mode"
	@cargo release --dry-run
	@cargo release --sign
	@TARGET_CC=x86_64-linux-musl-gcc cargo build --release --target=x86_64-unknown-linux-musl


build.docker:$(TAG)	## Build and push a docker container with binaries in release mode (takes a while)
	@echo "Build and release a docker container in release mode"

	@docker build -t "mariadb_test-$(TAG)" .
	@docker tag mariadb_test-$(TAG) ghcr.io/barkardk/integration_suite/mariadb_test:$(TAG)
    @docker push ghcr.io/barkardk/integration_suite/mariadb_test:$(TAG)

cross.compile:		## Setup rust for cross compilation
	@echo "Adding linker libraries"
	rustup target add x86_64-unknown-linux-musl
	@echo "Installing cross compiler"
	brew install FiloSottile/musl-cross/musl-cross




