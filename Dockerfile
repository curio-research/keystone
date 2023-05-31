# Build Geth in a stock Go builder container
FROM golang:1.20-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

# Get dependencies - will also be cached if we won't change go.mod/go.sum
COPY go.mod /go-ethereum/
COPY go.sum /go-ethereum/
RUN cd /go-ethereum && go mod download

ADD . /go-ethereum
RUN cd /go-ethereum && go run build/ci.go install -static ./cmd/geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

WORKDIR /curio-chain

RUN apk add --no-cache ca-certificates bash
COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/
COPY ./scripts/docker_run.sh ./
COPY ./pass-private.txt ./
COPY ./genesis_private.json ./
COPY ./data/keystore ./data/keystore

RUN chmod +x ./docker_run.sh

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["./docker_run.sh"]
