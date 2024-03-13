PATH := $(PATH):$(shell go env GOPATH)/bin

.PHONY: postgresinit
postgresinit:
	@echo "Create and start postgres database in docker"
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

.PHONY: dropdb
dropdb:
	@echo "Drop postgres database from docker"
	docker exec -it postgres dropdb postgres

.PHONY: test
test:
	@echo "Start tests"
	go test ./server/internal/service/command ./server/internal/handler/command

.PHONY: deps
deps:
	go mod tidy

.PHONY: swagger_init
swagger_init:
	@echo "Generate swagger gui"
	swag init --generalInfo  server/cmd/main.go -o server/docs

.PHONY: swag_ui
swag_ui:
	@echo "Open swagger index.html"
	open http://localhost:8000/swagger/index.html

.PHONY: server
server: deps swagger_init
	@echo "Running server"
	go run server/cmd/main.go

.PHONY: clean
clean:
	rm -rf ./build/scripts-launcher

.PHONY: build
build:
	docker-compose up