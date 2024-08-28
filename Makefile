
# Installs required software to work with the project
install:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest


format:
	gofmt -w .

lint:


# Generates database related and swagger documentation from request comments (also checks migration schema)
.PHONY: gen
gen:

