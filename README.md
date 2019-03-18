# GBDX Client

## master status

[![Build Status](https://travis-ci.org/afirth/fm-test.svg?branch=master)](https://travis-ci.org/afirth/fm-test) ![Docker Badge](https://img.shields.io/docker/cloud/build/afirth/fm-test.svg) ![Code Quality](https://goreportcard.com/badge/github.com/afirth/fm-test)

## godoc

- [![](https://godoc.org/github.com/afirth/fm-test?status.svg)](http://godoc.org/github.com/afirth/fm-test/api) api
- [![](https://godoc.org/github.com/afirth/fm-test?status.svg)](http://godoc.org/github.com/afirth/fm-test/gbdx) gbdx
- [![](https://godoc.org/github.com/afirth/fm-test?status.svg)](http://godoc.org/github.com/afirth/fm-test/transcode) transcode
- [![](https://godoc.org/github.com/afirth/fm-test?status.svg)](http://godoc.org/github.com/afirth/fm-test) main

## Author

@afirth 2019

# Installation

## Prerequisites

- Go >= 1.11
- If you want to use the makefile:
  - GNU Make (>3.81) (e.g. on OSX `brew install --with-default-names make`)
  - Bash 4.x

### Secrets

GBDX requires a username and password to get a token. For this project I built the token request into the server rather than as a separate service. _A valid username and password must be present in the environment._ This will be also be inherited by docker-compose and used to make the kubernetes secret.

```
export USERNAME=<gbdx-user-email>
export PASSWORD=<gbdx-password>
```

## Running locally

```
make up
```

## Running with docker-compose

```
make docker-compose-up
```

## Running in kubernetes

```
make kube-secret
make kube-up
```

### Note about images in kubernetes

By default, kube deployment pulls from dockerhub image. To use a local image:

```
eval $(minikube docker-env)
make docker-build
docker tag fm-test afirth/fm-test:0.0.1
```

# Building and testing

- `make test` runs unit tests
- `make test-e2e` runs unit and e2e tests
- `make build` builds for development
- `make build-final` builds for use in images, with vendored deps and no debugging

# Not Implemented

- retries for requests to GBDX
- rate limits for incoming requests (I prefer at gateway)
- decent json logging
- exhaustive tests

# Tested with

- OSX 10.14.2 (18C54)
- Minikube v0.35.0
- Kubectl v1.11.7
- Kubernetes v1.13.4
- Docker Client: Docker Engine - Community 18.09.1
- Docker Server: 18.06.2-ce
- GNU Make 3.81
- GNU bash, version 4.4.23
