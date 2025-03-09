# Variables
IMAGE_NAME = jicowan/tasky
TAG = latest

# Phony targets
.PHONY: build push all

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME):$(TAG) .

# Push the image to DockerHub
push: build
	docker login
	docker push $(IMAGE_NAME):$(TAG)

# Build and push
all: build push

# Help target
help:
	@echo "Available targets:"
	@echo "  build  - Build the Docker image"
	@echo "  push   - Push the image to DockerHub (requires login)"
	@echo "  all    - Build and push the image"
