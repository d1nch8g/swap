package main

import (
	"context"
	"errors"

	"github.com/jessevdk/go-flags"
	"ion.lc/d1nhc8g/bitchange/bestchange"
	"ion.lc/d1nhc8g/bitchange/gen/database"
	"ion.lc/d1nhc8g/bitchange/server"

	"github.com/jackc/pgx/v5"

	"github.com/labstack/echo/v4"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var opts struct {
	Port       string `long:"port" env:"PORT"`
	Database   string `long:"database" env:"DATABASE"`
	Bestchange string `long:"bestchange" env:"BESTCHANGE"`
	Dir        string `long:"dir" env:"DIR" default:"dist"`
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

	sqlc := database.New(conn)
	bestchange := bestchange.New(opts.Bestchange)
	echo := echo.New()

	server.Run(opts.Dir, opts.Port, echo, sqlc, bestchange)
}
