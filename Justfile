set positional-arguments := true
set dotenv-load := true
set shell := ["bash", "-c"]

# List all tasks
_default:
    @just --list

# Docker login
login:
    @echo "Logging into Docker Hub..."
    @docker login --username "${DOCKER_USERNAME}" --password "${DOCKER_PASSWORD}"
    @echo "Docker login complete."

# Build Docker image FOR amd64 and arm64
build: login
    @echo "Building Docker image..."
    @docker buildx rm proxytunnel || true
    @docker buildx create --name proxytunnel --use
    @docker buildx build \
        --platform linux/amd64,linux/arm64 \
        --push \
        -t "${DOCKER_USERNAME}/proxytunnel:alpine-go" .
    @echo "Docker image built and pushed."
    @docker buildx rm proxytunnel || true