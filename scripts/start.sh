#!/usr/bin/env bash
set -euo pipefail

echo "Building and starting ai-context via docker compose..."
docker compose up -d --build

echo "Service started. Access: http://localhost:8080"
