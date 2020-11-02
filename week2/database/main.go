package main

import (
	"log"
	"os"

	"github.com/hieu/golang-training/week2/command"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			&command.InfoCommand,
			&command.ServeCommand,
			&command.MigrateCommand,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
