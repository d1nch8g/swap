package main

import (
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Query  bool `short:"Q" long:"query"`
	Remove bool `short:"R" long:"remove"`
	Sync   bool `short:"S" long:"sync"`
	Push   bool `short:"P" long:"push"`
	Build  bool `short:"B" long:"build"`
}

func main() error {
	_, err := flags.NewParser(&opts, flags.IgnoreUnknown).Parse()
	if err != nil {
		return err
	}

}
