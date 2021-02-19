DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST)  | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

NON_VENDOR_PACKAGES=$(shell go list ./... | grep -v "vendor")
ifndef TAG
  TAG=$(shell git  rev-parse --short @)
endif
DOCKER_REGISTRY=ghcr.io/barkardk
fmt:
	go fmt $(NON_VENDOR_PACKAGES)
vet:
	go vet $(NON_VENDOR_PACKAGES)

restart: docker.stop docker.start ## Restart docker-compose
build: fmt vet build.linux build.docker.and.push ## Build a linux binary in release mode , build a docker image and push to registry
start: docker.start ## Start docker-compose
stop: docker.stop ## Stop docker-compose
darwin: build.darwin ## Build a dev binary on mac os
TAG: ## Set the env variable TAG to override the default git rev-parse


.PHONY: docker.start
docker.start:
	docker-compose up -d --remove-orphans;

.PHONY: docker.stop
docker.stop:
	docker-compose stop;

.PHONY: docker.build
docker.build:
	docker build -t rabbitmq-client

.PHONY: test
test:
	go test -tags=rabbitintegration  it/mq_test.gp -v -count=1;

.PHONY: build.linux
build.linux:
	mkdir -p target/linux
	env GO111MODULE=on CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o target/linux/mq_test -tags prod -mod=readonly -ldflags "-s -w " ./it
	chmod 755 target/linux/mq_test

.PHONY: build.docker.and.push
build.docker.and.push:
	docker build -t ${TAG} .
	docker tag ${TAG} ${DOCKER_REGISTRY}/integration_suite/rabbitmq-client:${TAG}
	docker push ${DOCKER_REGISTRY}/integration_suite/rabbitmq-client:${TAG}

.PHONY: build.darwin
build.darwin:
	mkdir -p target/darwin
	go build -o target/darwin/mq_test -tags dev it/main.go
	chmod 755 target/darwin/mq_test





