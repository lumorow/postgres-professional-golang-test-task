PATH := $(PATH):$(shell go env GOPATH)/bin

.PHONY: postgresinit
postgresinit:
	@echo "Create and start postgres database in docker"
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

.PHONY: dropdb
dropdb:
	@echo "Drop postgres database from docker"
	docker exec -it postgres dropdb postgres

.PHONY: migrationup
migrationup:
	@echo "Migration tables"
	migrate -path server/db/migrations -database "postgresql://root:password@localhost:5432/postgres?sslmode=disable" -verbose up

.PHONY: migrationdown
migrationdown:
	@echo "Delete tables"
	migrate -path server/db/migrations -database "postgresql://root:password@localhost:5432/postgres?sslmode=disable" -verbose down

.PHONY: test
test:
	go test ./server/internal/...

.PHONY: deps
deps:
	go mod tidy

.PHONY: swagger_init
swagger_init:
	@echo "Generate swagger gui"
	swag init --generalInfo  server/cmd/main.go -o server/docs

.PHONY: swag_gui
swag_gui:
	@echo "Open swagger index.html"
   	open http://localhost:8000/swagger/index.html

.PHONY: server
server: deps swagger_init
	@echo "Running server"
	go run server/cmd/main.go

.PHONY: build
build:
	docker-compose up
