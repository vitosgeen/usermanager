include ./configs/.env
build:
	go build ./cmd/usermanager/main.go

run:
	go run ./cmd/usermanager/main.go

run_grpc_server:
	go run ./cmd/usermanager/grpc/server/main.go

run_grpc_client:
	go run ./cmd/usermanager/grpc/client/main.go

test:
	go test -v -cover ./...

run-linter:
	echo "Starting linters"
	golangci-lint run ./...

# ==============================================================================
# Docker compose commands

dev:
	echo "Starting docker dev environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.dev.yml up usermanager --build

prod:
	echo "Starting docker prod environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.prod.yml up usermanager --build

local:
	echo "Starting docker local environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.local.yml up usermanager --build

# ==============================================================================
# Go migrate postgresql	
force:
	migrate -path ./db/migrations -database force 1 'postgres://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST_LOCAL}:${POSTGRES_PORT_LOCAL}/${POSTGRES_DBNAME}?sslmode=disable' -verbose force

version:
	migrate -path ./db/migrations -database version 1 'postgres://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST_LOCAL}:${POSTGRES_PORT_LOCAL}/${POSTGRES_DBNAME}?sslmode=disable' -verbose version

migrate_up:
	migrate -path ./db/migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST_LOCAL}:${POSTGRES_PORT_LOCAL}/${POSTGRES_DBNAME}?sslmode=disable' -verbose up

migrate_down:
	migrate -path ./db/migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST_LOCAL}:${POSTGRES_PORT_LOCAL}/${POSTGRES_DBNAME}?sslmode=disable' -verbose down

migrate_down_version:
	migrate -path ./db/migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST_LOCAL}:${POSTGRES_PORT_LOCAL}/${POSTGRES_DBNAME}?sslmode=disable' goto $(VERSION)

migratecreate:
	migrate create -ext sql -dir ./db/migrations $(MNAME)