package router

import "alta-test/controller/handler"

func init() {
	handlers["users"] = &handler.UserHandler{}
}
