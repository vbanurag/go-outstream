# syntax=docker/dockerfile:1

FROM golang:1.20 AS build-stage

WORKDIR /app

COPY . .
RUN go build -o microservice

FROM debian:buster
WORKDIR /app
COPY --from=build-stage /app/config.json .
COPY --from=build-stage /app/microservice .

CMD ["./microservice"]
