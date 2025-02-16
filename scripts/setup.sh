#!/bin/bash
# Setup development environment
go mod init $PROJECT_NAME
go mod tidy

# Install necessary tools
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
