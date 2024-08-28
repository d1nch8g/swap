package main

import (
	"context"

	"github.com/jessevdk/go-flags"

	"github.com/jackc/pgx/v5"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

var opts struct {
	Port     int32  `long:"port" env:"PORT"`
	Database string `long:"database" env:"DATABASE"`
}

func main() {
	_, err := flags.NewParser(&opts, flags.IgnoreUnknown).Parse()
	if err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(context.Background(), opts.Database)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	m, err := migrate.New(
		"file://db/migrations",
		opts.Database,
	)
	if err != nil {
		panic(err)
	}
	m.Steps(2)

}
