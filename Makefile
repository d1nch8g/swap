
# Installs required software to work with the project
install:
	go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/swaggo/swag/cmd/swag@latest

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
	swag init -o . --ot yaml 

watch:
	for number in 1 2 3 4 ; do \
        echo $$number ; \
    done
# swagger-codegen generate -i swagger.yaml -l javascript --disable-examples --flatten-inline-schema -o gen/client

# ip on local network
ip:
	ip route get 1.2.3.4 | grep -oP '(?<=src )\S+' 