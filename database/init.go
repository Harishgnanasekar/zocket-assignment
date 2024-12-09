package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=" + viper.GetString("DB_HOST") + " user=" + viper.GetString("DB_USER") +
		" password=" + viper.GetString("DB_PASSWORD") + " dbname=" + viper.GetString("DB_NAME") +
		" port=" + viper.GetString("DB_PORT") + " sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Database connected")
}
