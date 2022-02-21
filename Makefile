-include ./config/.env.dev
-include ./config/.env.prod

MIGRATE := migrate -path ./db/migrations -database '${PSQL_CONN_STRING}'

default: 
	@echo "No command provided"

migrate:
	@echo "Running all migrations..."
	$(MIGRATE) up

migrate-drop:
	@echo "Dropping database schema..."
	$(MIGRATE) drop

migrate-reset:
	@echo "Resetting database..."
	make migrate-drop
	make migrate

migrate-reset-seed:
	@echo "Resetting database..."
	make migrate-drop
	make migrate
	@echo "Seeding database..."
	make seed-db

migrate-new:
	@read -p "Input name of new migration: " NAME; \
	migrate create -dir ./db/migrations -ext sql -tz UTC $$NAME

migrate-down:
	$(MIGRATE) down 1

seed-db:
	@psql -f ./db/seeding/default.sql '${PSQL_CONN_STRING}'

run-dev:
	@air -c ./config/.air.toml

format:
	@go fmt core/...

test-queries:
	@go test ./internal/entity/queries -v

test:
	@go test core... -v