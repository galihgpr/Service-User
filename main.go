package main

import (
	"alta-test/config"
	"fmt"
)

func main() {
	appConfig := config.Get()
	postgreDB := config.InitDB(appConfig)
	fmt.Println(postgreDB)
}
