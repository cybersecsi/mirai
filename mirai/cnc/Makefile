VERSION=$(shell cat VERSION.txt)
IMAGE_NAME=secsi/mirai-cnc

docker-build:
	docker build -t $(IMAGE_NAME):$(VERSION) .


docker-push:
	@echo "Building and pushing Docker image for arm64 and amd64 platforms with version: $(VERSION)"
	@docker buildx create --use --name mybuilder --bootstrap || echo "Builder already exists"
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--tag $(IMAGE_NAME):$(VERSION) \
		--tag $(IMAGE_NAME):latest \
		--push .


