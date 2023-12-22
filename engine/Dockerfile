FROM golang:1.21.0-alpine
RUN mkdir -p /src/github.com/gavr-games/reborn-mmorpg
WORKDIR /src/github.com/gavr-games/reborn-mmorpg

RUN go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8080
EXPOSE 8081
EXPOSE 8082
