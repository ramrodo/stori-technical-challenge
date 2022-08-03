# syntax=docker/dockerfile:1

## Build

FROM golang:1.18-alpine AS build
ENV GO111MODULE=on

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go ./
COPY models/ models/
RUN go build -o . ./...

## Deploy

FROM alpine

WORKDIR /

COPY --from=build /app/stori-technical-challenge ./stori-technical-challenge
COPY txns.csv ./
COPY email-template.html ./

ENTRYPOINT [ "./stori-technical-challenge" ]
