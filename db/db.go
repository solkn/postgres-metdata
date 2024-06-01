package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"postgres-metadata/utils"
	
)

var DB *gorm.DB

func ConnectDatabase(host string) *gorm.DB {
	dbConn, err := gorm.Open(postgres.Open(host), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	DB = dbConn
	if err != nil {
		panic(err)
	}


	utils.Logger.Println("::successfully connected to DB::")
	return dbConn

}
