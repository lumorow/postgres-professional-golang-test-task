postgresinit:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

postgres:
	docker exec -it postgres psql

createdb:
	docker exec -it postgres createdb --username=root --owner=root go_pstgs_commands

dropdb:
	docker exec -it postgres dropdb go-chat

migrationup:
	migrate -path server/db/migrations -database "postgresql://root:password@localhost:5432/postgres?sslmode=disable" -verbose up

migrationdown:
	migrate -path server/db/migrations -database "postgresql://root:password@localhost:5432/postgres?sslmode=disable" -verbose down

server: postgresinit migrationup
	go run server/cmd/main.go

.PHONY: postgresinit postgres createdb dropdb migrationup migrationdown