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

	AwsS3 struct {
		Bucket    string
		Region    string
		AccessKey string
		SecretKey string
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
	config.PostgreDB.Password = GetEnv("DB_PASSWORD", "kurosaki")
	config.PostgreDB.DBName = GetEnv("DB_NAME", "user")
	config.PostgreDB.SslMode = GetEnv("DB_SSL", "disable")
	config.PostgreDB.TimeZone = GetEnv("DB_TIMEZONE", "Asia/Jakarta")
	// Config S3
	config.AwsS3.Bucket = GetEnv("AWS_S3_BUCKET", "belajar-be")
	config.AwsS3.Region = GetEnv("AWS_S3_REGION", "ap-southeast-1")
	config.AwsS3.AccessKey = GetEnv("AWS_S3_ACCESS_KEY", "AKIAWEUURQ3T37U2Y2GF")
	config.AwsS3.SecretKey = GetEnv("AWS_S3_SECRET_KEY", "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")

	return &config
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		return value
	}

	return fallback
}
