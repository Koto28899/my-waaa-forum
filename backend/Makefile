postgres:
	docker run --name mywaaa -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:latest

createdb:
	docker exec -it mywaaa createdb --username=root --owner=root mywaaa

dropdb:
	docker exec -it mywaaa dropdb mywaaa

diveindb:
	docker exec -it mywaaa psql -d mywaaa -U root

migrateup:
	migrate -path db/migrations -database "postgres://root:1234@localhost:5432/mywaaa?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgres://root:1234@localhost:5432/mywaaa?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

ent:
	go generate ./ent

dev:
	go run main.go

build:
	docker build -t waaa:lastest .

.PHONY: postgres createdb dropdb migrateup migratedown test ent dev build