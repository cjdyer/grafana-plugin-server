.PHONY: all
all: build backend backend-build backend-clean frontend frontend-build frontend-clean build clean run

build: backend-build frontend-build
clean: backend-clean frontend-clean
run:
	@echo ">>> Starting backend and frontend (dev mode)..."
	@$(MAKE) -j2 backend frontend

backend:
	@echo ">>> Running Go backend..."
	go run pkg/main.go

backend-build:
	@echo ">>> Building Go backend..."
	go build -o build/server pkg/main.go

backend-clean:
	@echo ">>> Cleaning backend build..."
	rm -f build

frontend:
	@echo ">>> Starting frontend dev server..."
	npm run dev

frontend-build:
	@echo ">>> Building frontend bundle..."
	npm run build

frontend-clean:
	@echo ">>> Cleaning frontend dist..."
	npm run clean