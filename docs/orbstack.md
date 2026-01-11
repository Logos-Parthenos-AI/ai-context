Orbstack / Docker Compose

This repository includes a docker-compose.yml to run ai-context as a container suitable for Orbstack environments.

Quick commands (from repo root):

- ./scripts/start.sh   # Build image and start container
- ./scripts/stop.sh    # Stop and remove container
- ./scripts/status.sh  # Show container status and recent logs

The service will be exposed on port 8080 and is reachable at http://localhost:8080

Healthcheck is configured in docker-compose.yml and the image is built from the included Dockerfile.

If you need Orbstack-specific networking, map ports in the compose file or use Orbstack UI to expose the container.
