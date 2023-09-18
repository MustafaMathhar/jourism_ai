include .env
export
create:
	@read -p "Enter the migration: " INP; \
	migrate create -ext sql -dir datastore/migrations -seq $$INP
migrate:
	 migrate -database "mysql://${DSN}" -path=datastore/migrations up
force: 
	@read -p "Enter the version: " VER; \
	 migrate -database "mysql://${DSN}" -path=datastore/migrations force $$VER
build:
	go build -o bin/jourism_ai -ldflags="-s -w" main.go
test: build
	go run .
