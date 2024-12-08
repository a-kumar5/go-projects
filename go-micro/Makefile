FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp

## up: starts all container in the backgroup without forcing build

up:
		@echo "Starting Docker images..."
		docker compose up -d
		@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker
		@echo "Stoping docker images (if running)"
		docker compose down
		@echo "Build (when required) and starting docker images..."
		docker compose up --build -d
		@echo "Docker images build and started!"

## down: stop docker compose
down:
		@echo "Stopping docker compose..."
		docker compose down
		@echo "Done!"

## build_front: build the front end binary
build_front:
		@echo "Building front end library"
		cd ./front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
		@echo "Done!"

## start: start the front end
start: build_front
		@echo "Starting front end"
		cd ./front-end && ./${FRONT_END_BINARY} &

## stop: stop the front end
stop:
		@echo "Stopping front end..."
		@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
		@echo "Stopped front end!"