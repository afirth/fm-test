FROM golang:1.12-alpine AS builder

# generic golang builder, takes repo_name as a build arg
# builds using go mod and vendored deps, so use CI!
# afirth 2018 2019

RUN apk update && apk add --no-cache \
  make \
  git \
  ca-certificates \
  && update-ca-certificates

COPY . /src/app
RUN cd /src/app && \
      make build-final BINPATH=/go/bin/app SHELL=/bin/sh

# although we could use scratch, alpine makes debugging much easier and comes with (some) tls support
FROM alpine:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/app /go/bin/app

# Create don't run as god
RUN adduser -D -g '' appuser
USER appuser

ENTRYPOINT ["/go/bin/app"]
