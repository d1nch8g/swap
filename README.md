# Swap - open currency exchange spot software

This projects provides access to self hosted exchange spot which can be configured under your personal needs.

## Locally from repository

1. In repository build npm source files into `dist`.

```sh
npm run build
```

2. Build go source binaries (you have to have go installed on your machine):

```sh
go install github.com/go-bindata/go-bindata/go-bindata@latest
go-bindata -pkg web -o gen/web/web.go -fs -prefix "dist/" dist/...
go-bindata -pkg migr -o gen/migr/migr.go -fs -prefix "db/migrations/" db/migrations/...
```

3. Run project locally (with database in compose):

```sh
docker compose up
go run . --port 8080 --admin "email@example.com:password" --database "postgresql://user:password@localhost:5432/db?sslmode=disable" --telegram "@test" --email-creds "mail@example.com:password"
```

## Run as go CLI

Installation:

```
go install github.com/d1nch8g/swap@latest
```

Configuration flags:

```sh
--port             - Port on which to run application on
--host             - Hostname, should be inintialized on production runs
--database         - database connection string
--bestchange-token - token to access bestchange API
--admin            - admin creds "email:password"
--email-addr       - email client address
--email-port       - email client port
--email-creds      - email "login:password"
--telegram         - telegram link
--cert-file        - Cert file path (should be used for TLS)
--key-file         - Key file path (should be used for TLS)
-h --help          - Show this help message and exit
```

Minimal example (requires active postgresql database, can be started):

```sh
swap --port
```

## Run with docker

Image:

Configuration:

- `PORT` - port on which program will run
- `DATABASE` - database connection string
- `BESTCHANGE_TOKEN` - token to call bestchange api to receive exchanger rates info
- `ADMIN` - user:password of main admin account
- `API_ADDR` - link to api should be self referenced api address
- `EMAIL_ADDRESS` - host of email address
- `EMAIL_CREDS` - creds for email
- `TELEGRAM` - telegram link @tlg
- `CERT_FILE` - certificate file
- `KEY_FILE` - key file

<!--

http://192.168.0.105:8080/?currin=TON&currout=SBPRUB

add transactions to execute and create order functions

add validate card page

-->
