BALANCE_SERVICE_BINARY=binary_file/balanceServiceApp

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_tech_task
	@echo "stopping docker images (if running...)"
	docker-compose down
	@echo "building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "docker images built and started"

## down: stop docker compose
down:
	@echo "stopping docker compose..."
	docker-compose down
	@echo "done"

## tech_task: builds the tech task binary as a linux executable
build_tech_task:
	@echo "building tech task binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o ${BALANCE_SERVICE_BINARY} ./cmd/service
	@echo "done"

#Запуск интеграционных тестов
test.integration:
	go test -tags=integration ./integration_tests -v 

# Запуск UNIT тестов 
test.unit:
	go test ./...
