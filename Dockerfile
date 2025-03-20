# syntax=docker/dockerfile:1
ARG alpine_version=3.20
ARG golang_version=1.23

FROM golang:${golang_version}-alpine${alpine_version} AS builder

RUN apk add --no-cache \
                git \
                make

# Download dependencies first to cache them
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy the source code and build
WORKDIR /app/code
COPY . .
RUN make build

FROM scratch AS deploy
COPY --from=builder /app/code/bin/docker-dns-sync docker-dns-sync
# this pulls directly from the upstream image, which already has ca-certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/docker-dns-sync"]
