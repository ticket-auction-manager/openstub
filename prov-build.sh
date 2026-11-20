#!/bin/sh

pnpm --dir frontend build
go build -o build/openstub -ldflags="-w -s" ./cmd
