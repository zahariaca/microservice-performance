GO_REST=goRestApp

build_go_be:
	@echo "Building go rest microservice..."
	cd ../go-rest/ && env GOOS=linux CGO_ENABLED=0 go build -o ${GO_REST} ./cmd/api
	@echo "Done!"

build_spring_boot_rest:
	@echo "Building spring boot rest application"
	cd ../spring-boot-rest/ && ./gradlew bootBuildImage
	@echo "Done!"

build_spring_boot_webflux:
	@echo "Building spring boot webflux application"
	cd ../spring-boot-webflux/ && ./gradlew bootBuildImage
	@echo "Done!"

build_spring_boot_webflux_bad:
	@echo "Building spring boot webflux application"
	cd ../spring-boot-webflux-bad/ && ./gradlew bootBuildImage
	@echo "Done!"


## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build_go: build_go_be
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

up_build_all: build_go_be build_spring_boot_rest build_spring_boot_webflux build_spring_boot_webflux_bad
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

up_go: build_go_be
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d postgres go-rest

up_no_vt: build_spring_boot_rest
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d postgres sb-no-vt

up_vt: build_spring_boot_rest
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d postgres sb-vt

up_webflux: build_spring_boot_webflux
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d postgres sb-webflux

up_webflux_bad: build_spring_boot_webflux_bad
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d postgres sb-webflux-bad