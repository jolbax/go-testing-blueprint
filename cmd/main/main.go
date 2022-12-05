package main

import (
	"go-testing-blueprint/internal/controller"
	"go-testing-blueprint/internal/router"
)

func main() {
	ginEngine := router.NewRouter(controller.NewHTTP())
	_ = ginEngine.Run()
}
