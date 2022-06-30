FROM golang:1.18 AS build
WORKDIR /workspace
COPY . .
RUN LEDGER_ENABLED=false make build

FROM scratch
COPY --from=build /workspace/build/elestod /elestod

ENTRYPOINT ["/elestod"]
