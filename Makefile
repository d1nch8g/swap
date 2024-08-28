
# Installs required software to work with the project
install:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3

format:
	gofmt -w .

rundb:
	docker run --rm -it 

# Generates database related and swagger documentation from request comments (also checks migration schema)
.PHONY: gen
gen:

