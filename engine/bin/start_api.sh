#!/bin/sh

go mod download
# dlv debug ./cmd/api/
go run cmd/api/main.go
