APP_NAME := chat-backend
IMAGE_NAME := $(APP_NAME):local
CONTAINER_NAME := $(APP_NAME)
PORT := 8080

.PHONY: docker-build docker-run docker-stop compose-up compose-up-d compose-down compose-down-v compose-logs run stop

docker-build:
	docker build -t $(IMAGE_NAME) .

docker-run: docker-build
	docker run --rm -d \
		--name $(CONTAINER_NAME) \
		-p $(PORT):$(PORT) \
		$(IMAGE_NAME)

docker-stop:
	docker stop $(CONTAINER_NAME)

compose-up:
	docker compose up --build

compose-up-d:
	docker compose up --build -d

compose-down:
	docker compose down

compose-down-v:
	docker compose down -v

compose-logs:
	docker compose logs -f

run: docker-run

stop: docker-stop
