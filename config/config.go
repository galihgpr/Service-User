package config

import (
	"fmt"
	"os"
)

type AppConfig struct {
	PostgreDB struct {
		Username string
		Password string
		Host     string
		Port     string
		DBName   string
		TimeZone string
		SslMode  string
	}
}

var appConfig *AppConfig

func Get() *AppConfig {

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {
	var config AppConfig

	// Config DB
	config.PostgreDB.Host = GetEnv("DB_HOST", "localhost")
	config.PostgreDB.Port = GetEnv("DB_PORT", "5432")
	config.PostgreDB.Username = GetEnv("DB_USERNAME", "root")
	config.PostgreDB.Password = GetEnv("DB_PASSWORD", "mysecretnumber")
	config.PostgreDB.DBName = GetEnv("DB_NAME", "postgres")
	config.PostgreDB.SslMode = GetEnv("DB_SSL", "disable")
	config.PostgreDB.TimeZone = GetEnv("DB_TIMEZONE", "Asia/Jakarta")

	return &config
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		return value
	}

	return fallback
}
