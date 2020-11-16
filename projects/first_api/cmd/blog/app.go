package main

import (
	"log"
	"os"

	"github.com/hhhieu/golang-training/first_api/internal/pkg/command"
	"github.com/urfave/cli/v2"
)

var commander = command.NewCommander()

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "interval",
				Aliases: []string{"i"},
				Value:   10,
				Usage:   "Retry interval",
			},
			&cli.IntFlag{
				Name:    "retry-time",
				Aliases: []string{"r"},
				Value:   6,
				Usage:   "Retry time",
			},
			&cli.IntFlag{
				Name:    "start-period",
				Aliases: []string{"a"},
				Value:   20,
				Usage:   "Initialization time",
			},
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name:  "wait-database",
				Usage: "Waiting for database ready",
				Action: func(c *cli.Context) error {
					log.Printf("Start waiting database")
					// Initialization time
					startPeriod := c.Int("start-period")
					retryTime := c.Int("retry-time")
					interval := c.Int("interval")
					return commander.WaitForDatabase(startPeriod, retryTime, interval)
				},
			},
			&cli.Command{
				Name:  "serve",
				Usage: "Run the service",
				Action: func(c *cli.Context) error {
					return commander.Serve()
				},
			},
		},
	}
	app.Run(os.Args)
}
