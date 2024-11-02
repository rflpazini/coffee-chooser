# syntax=docker/dockerfile:1
FROM golang:1.23-alpine as base
LABEL authors="rflpazini@gmail.com"

RUN apk add --no-cache curl
RUN apk add --no-cache git ca-certificates openssh-client
RUN apk add --no-cache g++ && apk --no-cache add make

FROM base as builder
WORKDIR /app

COPY go.mod go.sum Makefile ./

RUN --mount=type=cache,target=/root/.cache \
    --mount=type=cache,target=/go/pkg/mod \
     make deps

COPY . ./
RUN make build

FROM alpine:3.20
WORKDIR /app

ARG APP_VERSION
ARG MONGODB_URL
ARG BRANCH_NAME

ENV APP_VERSION ${APP_VERSION}
ENV MONGODB_URL ${MONGODB_URL}
ENV BRANCH_NAME ${BRANCH_NAME}

COPY --from=builder /app/config/ ./config/
COPY --from=builder /app/scripts/ ./scripts/
COPY --from=builder /app/bin/coffee-chooser .

EXPOSE 8080
EXPOSE 443

ENTRYPOINT ["/app/coffee-chooser"]

