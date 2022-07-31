# syntax=docker/dockerfile:1

## Build

FROM golang:1.18-alpine AS build
ENV GO111MODULE=on

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY main.go ./
RUN go build -o /stori-app

## Deploy

FROM scratch

WORKDIR /

COPY --from=build /stori-app /stori-app
COPY txns.csv ./

ENTRYPOINT [ "/stori-app" ]
