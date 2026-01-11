#!/usr/bin/env bash
set -euo pipefail

# Use a project-local module cache to avoid permission issues creating the global cache
CACHE_DIR="$(pwd)/.gomodcache"
mkdir -p "$CACHE_DIR"
echo "Using GOMODCACHE=$CACHE_DIR"
GOMODCACHE="$CACHE_DIR" go build .
