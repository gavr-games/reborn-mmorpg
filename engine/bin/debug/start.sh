#!/bin/sh

go mod download
dlv debug ./cmd/server/
