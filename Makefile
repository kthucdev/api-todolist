IMAGE_NAME=go-todolist-api
CONTAINER_NAME=go-todolist-container
PORT=8080

# Build the application
all: build test

# Build Docker image
build:
	docker build -t $(IMAGE_NAME) .

# Run the application
run: 
	docker run -d -p 8080:8080 -v $(PWD):/app --name go-todolist-container go-todolist-api

# Stop container
stop:
	docker stop $(CONTAINER_NAME)

# Log the container
logs:
	docker logs -f $(CONTAINER_NAME)

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

clean: stop
	docker rm $(CONTAINER_NAME) || true
	docker rmi $(IMAGE_NAME) || true

.PHONY: build run clean