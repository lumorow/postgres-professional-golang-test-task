postgresinit:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

dropdb:
	docker exec -it postgres dropdb postgres

migrationup:
	migrate -path server/db/migrations -database "postgresql://root:password@localhost:5432/postgres?sslmode=disable" -verbose up

migrationdown:
	migrate -path server/db/migrations -database "postgresql://root:password@localhost:5432/postgres?sslmode=disable" -verbose down

server:
	go run server/cmd/main.go

go_swagger:
	cd server && swag init --generalInfo  server/cmd/main.go -o server/docs


.PHONY: postgresinit dropdb migrationup migrationdown server