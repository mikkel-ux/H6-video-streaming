package config

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func ConnectDB() {
	once.Do(func() {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf(
			"server=%s;port=%s;database=%s;trustServerCertificate=true;encrypt=disable",
			host, port, dbName,
		)

		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to database!")
		}
		DB = db
	})
}
