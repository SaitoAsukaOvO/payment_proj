#!/bin/bash

sudo rm -rf swagger-generated

set -e

go test ./...
go test ./... -json > test-report.json
go test ./... -coverprofile=coverage.out
