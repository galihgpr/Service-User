package config

import (
	"alta-test/entities"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.PostgreDB.Host,
		config.PostgreDB.Username,
		config.PostgreDB.Password,
		config.PostgreDB.DBName,
		config.PostgreDB.Port,
		config.PostgreDB.SslMode,
		config.PostgreDB.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Warn()
		fmt.Println(err)
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(entities.User{})
}
