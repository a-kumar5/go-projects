postgres:
	docker run --name postgres -p 5433:5432 -e POSTGRES_PASSWORD=postgres -d postgres:13-alpine

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

runpgcli:
	docker exec -it postgres psql -U postgres

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/simple_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc