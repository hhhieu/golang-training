package provider

import (
	"log"

	"github.com/hieu/golang-training/week2/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RP struct {
	DB *gorm.DB
}

func MustBuildRP(conf *config.Config) *RP {
	log.Print(conf.MSQLDNS)
	sqlDB, err := gorm.Open(mysql.Open(conf.MSQLDNS), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &RP{
		DB: sqlDB,
	}
}
