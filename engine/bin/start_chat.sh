#!/bin/sh

go mod download
# dlv debug ./cmd/chat/
go run cmd/chat/main.go
