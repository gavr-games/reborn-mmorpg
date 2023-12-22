#!/bin/sh

go mod download
# dlv debug ./cmd/server/
go run cmd/server/main.go
