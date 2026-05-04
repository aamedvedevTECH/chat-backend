APP_NAME := chat-backend
IMAGE_NAME := $(APP_NAME):local
CONTAINER_NAME := $(APP_NAME)
PORT := 8080

.PHONY: docker-build docker-run docker-stop run stop

docker-build:
	docker build -t $(IMAGE_NAME) .

docker-run: docker-build
	docker run --rm -d \
		--name $(CONTAINER_NAME) \
		-p $(PORT):$(PORT) \
		$(IMAGE_NAME)

docker-stop:
	docker stop $(CONTAINER_NAME)

run: docker-run

stop: docker-stop
