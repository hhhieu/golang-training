package command

import (
	"fmt"
	"sync"

	"github.com/hieu/golang-training/week2/config"
	"github.com/hieu/golang-training/week2/provider"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var InfoCommand = cli.Command{
	Name:  "info",
	Usage: "run the service",
	Action: func(c *cli.Context) error {
		fmt.Println("Show info")
		return nil
	},
}

var ServeCommand = cli.Command{
	Name:  "serve",
	Usage: "run the service",
	Action: func(c *cli.Context) error {
		config := config.NewDefaultConfig()
		rp := provider.MustBuildRP(config)
		wg := sync.WaitGroup{}
		for j := 0; j < 5; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < 1000; i++ {
					rp.DB.Create(&Product{Code: fmt.Sprintf("%d", i), Price: 100})
				}
			}()
		}
		wg.Wait()
		return nil
	},
}
var MigrateCommand = cli.Command{
	Name:  "migrate",
	Usage: "run the service",
	Action: func(c *cli.Context) error {
		// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
		config := config.NewDefaultConfig()
		rp := provider.MustBuildRP(config)
		rp.DB.AutoMigrate(&Product{})
		return nil
	},
}
