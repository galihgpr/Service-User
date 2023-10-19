package main

import (
	"alta-test/controller/router"
)

func main() {

	e := router.Router()

	// Run Apps
	e.Run(":8000")
}
