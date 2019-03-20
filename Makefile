# @afirth 2018-11

.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --warn-undefined-variables
SHELL = /bin/bash
.SUFFIXES:

SVC_NAME ?= $(notdir $(CURDIR))
PORT ?= 8080
VERSION ?= $(shell cat VERSION)
#https://github.com/golang/go/issues/23439#issuecomment-433118300
BINPATH ?= $(CURDIR)/bin/app

all: build test docker-build

test:
	go test ./...

test-e2e: check-credentials
	GOTEST_E2E=1 go test -v -mod=vendor ./...

# builds ignoring vendor for local dev, then updates
build:
	go mod tidy
	go mod vendor
	GOMOD111=on go build -ldflags "-X main.Version=$(VERSION)" -o $(BINPATH)

build-final:
	GOOS=linux GOARCH=amd64 GO111MOD=on \
		go build \
			-mod=vendor \
			-ldflags "-w -s -X main.Version=$(VERSION)" \
			-o $(BINPATH)

up: check-credentials build
	HTTPADDR=:$(PORT) $(BINPATH)

docker-build:
	docker build -t $(SVC_NAME):$(VERSION) .

docker-up: check-credentials
	@docker run \
		-e USERNAME=$(USERNAME) \
		-e PASSWORD=$(PASSWORD) \
		-e HTTPADDR=':$(PORT)' \
		-p $(PORT):$(PORT) \
		--rm $(SVC_NAME):$(VERSION)
	@Service is running on localhost:$(PORT)

docker-compose-up: check-credentials
	@USERNAME=$(USERNAME) PASSWORD=$(PASSWORD) docker-compose up
	@Service is running on localhost:$(PORT)

kube-secret: check-credentials
	@kubectl create secret generic $(SVC_NAME)-secret \
--from-literal=username=$(USERNAME) \
--from-literal=password=$(PASSWORD) || echo "Secret exists"

kube-up:
	kubectl apply -f manifests/
	@(set +e kubectl config current-context | grep -n minikube && echo service will be available at: `minikube service $(SVC_NAME) --url`) || echo
	@echo 'if using port forwarding: kubectl port-forward svc/$(SVC_NAME) $(PORT):$(PORT)'
	
# ignore ./vendor with xargs until https://github.com/golang/lint/issues/320
lint:
	@go list ./... | xargs -L1 golint

check-credentials:
	@test $(USERNAME) || (echo "USERNAME is not set. Try export USERNAME=<gbdx-username>" && exit 60)
	@test $(PASSWORD) || (echo "PASSWORD is not set. Try export PASSWORD=<gbdx-password>" && exit 61)

.PHONY: all test build up docker-build docker-up docker-compose-up kube-secret kube-up lint check-credentials
