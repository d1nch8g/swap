package main

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"strings"

	"ion.lc/d1nhc8g/bitchange/bestchange"
	"ion.lc/d1nhc8g/bitchange/gen/database"
	"ion.lc/d1nhc8g/bitchange/server"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jessevdk/go-flags"
	"github.com/labstack/echo/v4"
)

var opts struct {
	Port       string `long:"port" env:"PORT" default:"8080"`
	Database   string `long:"database" env:"DATABASE" default:"postgresql://user:password@localhost:5432/db?sslmode=disable"`
	Dir        string `long:"dir" env:"DIR" default:"dist"`
	Bestchange string `long:"bestchange" env:"BESTCHANGE"`
	Tls        string `long:"tls" env:"TLS"`
	Admin      string `long:"admin" env:"ADMIN" default:"support@inswap.in:password"`
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

	conn, err := pgx.Connect(context.Background(), opts.Database)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	e := echo.New()
	sqlc := database.New(conn)
	bc := bestchange.New(opts.Bestchange)

	hasher := sha512.New()
	hasher.Write([]byte(strings.Split(opts.Admin, ":")[0]))
	passhash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	_, err = sqlc.CreateUser(context.Background(), database.CreateUserParams{
		Email:     strings.Split(opts.Admin, ":")[1],
		Verified:  true,
		Passwhash: passhash,
		Admin:     true,
		Active:    true,
		Busy:      false,
		Token:     "nil",
	})
	if err != nil && !strings.Contains(err.Error(), "duplicate key value violates unique constraint ") {
		panic(err)
	}

	server.Run(opts.Dir, opts.Port, opts.Tls, e, sqlc, bc)
}
