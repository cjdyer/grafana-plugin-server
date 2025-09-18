.PHONY: all clean build run docker-build docker-run

all: build

build: build-frontend build-backend

build-frontend:
	@echo ">>> Building frontend..."
	npm run build

build-backend:
	@echo ">>> Building Go server..."
	go build -o build/server pkg/main.go

clean:
	rm -rf build dist node_modules

run: build
	@echo ">>> Running server on :3838..."
	./build/server

build-docker:
	@echo ">>> Building Docker image..."
	docker build -t grafana-plugin-server:latest .

run-docker:
	@echo ">>> Running Docker container..."
	docker compose up --build
