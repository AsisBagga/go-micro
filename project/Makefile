FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp

## up: Starts all containers in the background without forcing build
up:
		@echo "Starting Docker images.."
		docker-compose up -d
		@echo "Docker image started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker
		@echo "Stopping docker images..(If running..)"
		docker-compose down
		@echo "Building and starting docker images.."
		docker-compose up --build -d
		@echo "Docker images build and started!"

## down: stop docker compose
down:
		@echo "Stopping docker containers..."
		docker-compose down
		@echo "Done!"

## build_broker: builds the broker binarry as a linux executable. This can be used by alpine inage. 
build_broker:
		@echo "Building broker binary.."
		cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
		@echo "Done!"

## build_front: build front end binary
build_front:
		@echo "Building front end binary..."
		cd ../front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
		@echo "Done!"

## start: starts the front end binary
start: 
		@echo "Starting front end binary"
		cd ../front-end && ./${FRONT_END_BINARY} &

## stops the front end Binary
stop:
		@echo "Stopping front end binary... "
		@-pkill -SIGTERM -f && ./${FRONT_END_BINARY}
		@echo "Stopped front end"