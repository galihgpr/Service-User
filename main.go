package main

import (
	"alta-test/config"
	"alta-test/model"
	"fmt"
)

func main() {
	appConfig := config.Get()
	postgreDB := config.InitDB(appConfig)
	fmt.Println(postgreDB)
	config.MigrateDB(postgreDB)

	// Inject Model
	modelUser := model.NewModelDB(postgreDB)
}
