postgresinit:
	docker run --name postgres -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:last-last

postgres:
	docker exec -it postgres psql

createdb:
	docker exec -it postgres createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it postgres dropdb go-chat

migrationup:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/chat?sslmode=disable" -verbose up

migrationdown:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/chat?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrationup migrationdown