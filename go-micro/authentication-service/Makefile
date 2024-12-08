c_m:
	# creates a new create_migration
	migrate create -ext sql -dir db/migrations -seq $(name)

db_up:
	docker exec -it micro_postgres createdb --username=root --owner=root micro_db

db_down:
	docker exec -it micro_postgres dropdb --username=root micro_db

migrate_up:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5433/micro_db?sslmode=disable" up

migrate_down:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5433/micro_db?sslmode=disable" down