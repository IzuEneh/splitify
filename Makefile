include .env
export

# Default environment variables
REDIRECT_URI ?= http://localhost:8080/playlists
CLIENT_ID ?= $(SPOTIFY_CLIENT_ID)
BUILD_COMMAND = build:dev
IMAGE_NAME ?= splitify-image
CURRENT_DATE = $(shell date '+%Y-%m-%d')
CONTAINER_NAME ?= splitify-dev

.PHONY: build run stop clean all help

# Build the Docker image
build: 
	docker build --build-arg BUILD_COMMAND=$(BUILD_COMMAND) --build-arg CLIENT_ID=$(CLIENT_ID) --build-arg REDIRECT_URI=$(REDIRECT_URI) -t $(IMAGE_NAME) .

# Run the container from the built image
run:
	docker run -d --name $(CONTAINER_NAME) -p 8080:8080 $(IMAGE_NAME)

# Stop and remove the running container
stop:
	docker stop $(CONTAINER_NAME) \
    docker rm $(CONTAINER_NAME)

# Remove the built image
clean:
	docker rmi $(IMAGE_NAME)

# Default target
all: build run

# Log environment variable value for testing
env:
	@echo "CLIENT_ID=$(SPOTIFY_CLIENT_ID)"