package main

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/d1nch8g/swap/bestchange"
	"github.com/d1nch8g/swap/email"
	"github.com/d1nch8g/swap/gen/database"
	"github.com/d1nch8g/swap/gen/migr"
	"github.com/d1nch8g/swap/server"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jessevdk/go-flags"
	"github.com/labstack/echo/v4"
)

//	@title			Swap exchange api
//	@version		1.0
//	@description	Swap exchange api.
//	@termsOfService	http://github.com/d1nch8g/swap

//	@contact.name	Swap Support
//	@contact.url	http://github.com/d1nch8g/swap
//	@contact.email	support@chx.su

//	@license.name	MIT
//	@license.url	https://github.com/d1nch8g/swap/src/branch/main/LICENSE

//	@host		localhost:8080
//	@BasePath	/api
//	@schemes	http https

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Token authorization for internal operations

var opts struct {
	Port            string `long:"port" env:"PORT" default:"8080"`
	Host            string `long:"host" env:"HOST"`
	Database        string `long:"database" env:"DATABASE" default:"postgresql://user:password@localhost:5432/db?sslmode=disable"`
	BestchangeToken string `long:"bestchange-token" env:"BESTCHANGE_TOKEN"`
	BestchangeLink  string `long:"bestchange-link" env:"BESTCHANGE_LINK" default:"https://bestchange.ru/"`
	Admin           string `long:"admin" env:"ADMIN" default:"support@chx.su:password"`
	EmailAddress    string `long:"email-addr" env:"EMAIL_ADDRESS" default:"mail.hosting.reg.ru"`
	EmailPort       int    `long:"email-port" env:"EMAIL_PORT" default:"587"`
	EmailCreds      string `long:"email-creds" env:"EMAIL_CREDS" default:"support@chx.su:password"`
	Telegram        string `long:"telegram" env:"TELEGRAM"`
	CertFile        string `long:"cert-file" env:"CERT_FILE"`
	KeyFile         string `long:"key-file" env:"KEY_FILE"`
	Help            bool   `short:"h" long:"help"`
}

func Help() string {
	return `Server parameters:
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
-h --help          - Show this help message and exit`
}

func main() {
	_, err := flags.NewParser(&opts, flags.IgnoreUnknown).Parse()
	if err != nil {
		panic(err)
	}
	if opts.Help {
		fmt.Println(Help())
		return
	}

	spew.Dump(opts)

	s := bindata.Resource(migr.AssetNames(),
		func(name string) ([]byte, error) {
			return migr.Asset(name)
		})

	d, err := bindata.WithInstance(s)
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithSourceInstance(
		"go-bindata",
		d,
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
		"https://"+opts.Host,
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

	server.Run(opts.Port, opts.Host, opts.CertFile, opts.KeyFile, strings.Split(opts.Admin, ":")[0], opts.Telegram, opts.BestchangeLink, e, conn, sqlc, bc, mail)
}
