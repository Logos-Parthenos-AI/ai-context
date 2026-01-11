#!/usr/bin/env bash
set -euo pipefail

echo "Stopping ai-context container..."
docker compose down

echo "Stopped."
