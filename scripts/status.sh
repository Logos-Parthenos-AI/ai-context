#!/usr/bin/env bash
set -euo pipefail

echo "Container status:"
docker compose ps

echo "Logs (last 200 lines):"
docker compose logs --no-color --tail=200 ai-context || true
