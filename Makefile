include .env
export

IMAGE_NAME = splitify-image
CURRENT_DIR = $(shell pwd)

# Docker build command
DOCKER_BUILD = docker build  --build-arg "GIN_MODE=$(GIN_MODE)" --build-arg "BUILD_COMMAND=build:dev" --build-arg "CLIENT_ID=$(SPOTIFY_CLIENT_ID)" --build-arg "REDIRECT_URI=$(REDIRECT_URI)" -t $(IMAGE_NAME) .

# Docker run command
DOCKER_RUN = docker run -it --rm -p 8080:8080 $(IMAGE_NAME)

LOCAL_BUILD_NODE = cd $(CURRENT_DIR)/client && VITE_REDIRECT_URI=$(REDIRECT_URI) VITE_CLIENT_ID=$(SPOTIFY_CLIENT_ID) npm run build:dev -- --outDir "../dev/static" --emptyOutDir 

LOCAL_BUILD_GO = cd $(CURRENT_DIR)/server && go build -o $(CURRENT_DIR)/dev/server .

RUN_SERVER = cd $(CURRENT_DIR)/dev && ./server

# Targets
.PHONY: prod clean local go docker

# Default target
local:
	$(LOCAL_BUILD_NODE)
	$(LOCAL_BUILD_GO)
	$(RUN_SERVER)

clean: clean-local
	docker rmi -f $(IMAGE_NAME) || true


clean-local:
	rm -rf dev/

go:
	$(LOCAL_BUILD_GO)
	$(RUN_SERVER)

node: 
	$(LOCAL_BUILD_NODE) --watch

docker:
	$(DOCKER_BUILD)
	$(DOCKER_RUN)
