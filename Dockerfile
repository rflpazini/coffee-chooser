FROM golang:1.21-alpine as base
LABEL authors="rflpazini"
RUN apk add --no-cache curl
RUN apk add --no-cache git ca-certificates openssh-client
RUN apk add --no-cache g++ && apk --no-cache add make

FROM base as builder
WORKDIR /go/src

COPY go.mod go.sum Makefile ./

RUN --mount=type=cache,target=/root/.cache \
    --mount=type=cache,target=/go/pkg/mod \
     make deps

COPY . ./
RUN make build

FROM alpine:3.18
WORKDIR /app

COPY --from=builder /go/src/config/ ./config/
COPY --from=builder /go/src/bin/coffee-chooser .

EXPOSE 8080

ENTRYPOINT ["/app/coffee-chooser"]

