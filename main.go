package main

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"strings"

	"ion.lc/d1nhc8g/inswap/bestchange"
	"ion.lc/d1nhc8g/inswap/email"
	"ion.lc/d1nhc8g/inswap/gen/database"
	"ion.lc/d1nhc8g/inswap/server"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jessevdk/go-flags"
	"github.com/labstack/echo/v4"
)

//	@title			Inswap exchange api
//	@version		1.0
//	@description	Inswap exchange api.
//	@termsOfService	http://github.com/d1nch8g/inswap

//	@contact.name	Inswap Support
//	@contact.url	http://github.com/d1nch8g/inswap
//	@contact.email	support@swagger.io

//	@license.name	MIT
//	@license.url	https://github.com/d1nch8g/inswap/src/branch/main/LICENSE

//	@host		localhost:8080
//	@BasePath	/api
//	@schemes	http https

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Token authorization for internal operations

var opts struct {
	Port            string `long:"port" env:"PORT" default:"8080"`
	Database        string `long:"database" env:"DATABASE" default:"postgresql://user:password@localhost:5432/db?sslmode=disable"`
	ServeDir        string `long:"serve-dir" env:"SERVE_DIR" default:"dist"`
	BestchangeToken string `long:"bestchange-token" env:"BESTCHANGE_TOKEN"`
	LetsEncryptAddr string `long:"tls" env:"TLS"`
	Admin           string `long:"admin" env:"ADMIN" default:"support@inswap.in:password"`
	ApiAddr         string `long:"api-addr" env:"API_ADDRESS" default:""`
	EmailAddress    string `long:"email-addr" env:"EMAIL_ADDRESS" default:"mail.hosting.reg.ru"`
	EmailPort       int    `long:"email-port" env:"EMAIL_PORT" default:"587"`
	EmailCreds      string `long:"email-creds" env:"EMAIL_CREDS" default:"support@inswap.in:password"`
}

func main() {
	_, err := flags.NewParser(&opts, flags.IgnoreUnknown).Parse()
	if err != nil {
		panic(err)
	}

	m, err := migrate.New(
		"file://db/migrations",
		opts.Database,
	)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			panic(err)
		}
	}

	conn, err := pgxpool.New(context.Background(), opts.Database)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	e := echo.New()
	sqlc := database.New(conn)
	bc := bestchange.New(opts.BestchangeToken)
	mail := email.New(
		opts.EmailAddress,
		strings.Split(opts.EmailCreds, ":")[0],
		strings.Split(opts.EmailCreds, ":")[1],
		opts.ApiAddr,
		opts.EmailPort,
	)

	hasher := sha512.New()
	hasher.Write([]byte(strings.Split(opts.Admin, ":")[1]))
	passhash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	_, err = sqlc.CreateUser(context.Background(), database.CreateUserParams{
		Email:     strings.Split(opts.Admin, ":")[0],
		Verified:  true,
		Passwhash: passhash,
		Admin:     true,
		Operator:  true,
		Busy:      false,
		Token:     uuid.New().String(),
	})
	if err != nil && !strings.Contains(err.Error(), "duplicate key value violates unique constraint ") {
		panic(err)
	}

	server.Run(opts.ServeDir, opts.Port, opts.LetsEncryptAddr, e, conn, sqlc, bc, mail)
}
