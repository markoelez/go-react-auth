FROM golang:1.14-alpine3.11

WORKDIR /backend

COPY . ./

RUN go mod download

WORKDIR /backend/cmd/api

RUN echo "building api dockerfile" && ls
