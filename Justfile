MIGRATE := "migrate -path ./db/migrations -database '${PSQL_CONN_STRING}'"
MIGRATE_TEST := "migrate -path ./db/migrations -database '${TEST_PSQL_CONN_STRING}'"

default: 
	echo No command provided

format:
	go fmt core/...

test:
	go test core... -v

run-dev:
	air -c ./config/.air.toml
	
seed-apidata:
	psql -f ./db/seeding/apidevdata.sql '${PSQL_CONN_STRING}'

migrate:
	echo "Running all migrations..."
	{{MIGRATE}} up

migrate-drop:
	echo "Dropping database schema..."
	{{MIGRATE}} drop

migrate-reset:
	echo "Resetting database..."
	make migrate-drop
	make migrate

migrate-new:
	read -p "Input name of new migration: " NAME; \
	migrate create -dir ./db/migrations -ext sql -tz UTC $$NAME

migrate-down:
	{{MIGRATE}} down 1

seed-querytestdata:
	echo Seeding database...
	psql -f ./db/seeding/querytestdata.sql '${TEST_PSQL_CONN_STRING}'

test-queries:
	echo Running tests...
	go test ./internal/entity/queries -v

test-queries-thorough:
	echo Resetting database...
	make migrate-drop-test-db
	make migrate-test-db
	make seed-querytestdata
	make test-queries

migrate-test-db:
	{{MIGRATE_TEST}} up

migrate-drop-test-db:
	{{MIGRATE_TEST}} drop
