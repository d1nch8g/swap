
# Installs required software to work with the project
install:
	go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3

format:
	gofmt -w .

rundb:
	docker run --rm -it 

migrate:
	migrate -path db/migrations/ -database "postgresql://user:password@localhost:5432/db?sslmode=disable" -verbose up

# Generates database related and swagger documentation from request comments (also checks migration schema)
.PHONY: gen
gen:
	sqlc generate

