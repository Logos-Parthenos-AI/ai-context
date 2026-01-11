Orbstack / Docker Compose

This repository includes a docker-compose.yml to run ai-context as a container suitable for Orbstack environments.

Quick commands (from repo root):

- ./scripts/start.sh   # Build image and start container
- ./scripts/stop.sh    # Stop and remove container
- ./scripts/status.sh  # Show container status and recent logs

The service will be exposed on port 8080 and is reachable at http://localhost:8080

Healthcheck is configured in docker-compose.yml and the image is built from the included Dockerfile.

If you need Orbstack-specific networking, map ports in the compose file or use Orbstack UI to expose the container.

Persistence

The docker-compose.yml sets restart: unless-stopped so the container will be restarted by Docker if it stops. To ensure the service comes back after a host reboot or Orbstack restart:

- Ensure Orbstack (or your Docker runtime) is configured to start on login.
- Confirm the container is running with: docker compose ps
- If needed, start it on boot with: docker compose up -d

Note: Orbstack provides UI options to manage auto-start; consult Orbstack docs for "start on login" settings.
