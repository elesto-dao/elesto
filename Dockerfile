FROM golang:1.18 AS build
WORKDIR /workspace
COPY . .
RUN LEDGER_ENABLED=false make build

FROM ghcr.io/glebiller/tendermint-readiness:0.1.0
COPY --from=build /workspace/build/elestod /elestod

ENTRYPOINT ["/elestod"]
