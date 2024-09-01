
# Installs required software to work with the project
install:
	go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	npm install

format:
	gofmt -w .

rundb:
	docker compose up

migrate:
	migrate -path db/migrations/ -database "postgresql://user:password@localhost:5432/db?sslmode=disable" -verbose up

run:
	go run . --port 8080 --database "postgresql://user:password@localhost:5432/db?sslmode=disable"

# Generates database related and swagger documentation from request comments (also checks migration schema)
.PHONY: gen
gen:
	sqlc generate
	swag fmt
	swag init -o . --ot yaml
