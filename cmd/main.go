package main

import (
	"fmt"
	"github.com/go-nag/configuration/cfgl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	manager, err := cfgl.LoadConfigFile("local")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	loggerEnabled, err := manager.Get("server.logging")
	if err != nil {
		log.Fatal(err)
	}

	if loggerEnabled == "enabled" {
		log.Println("Using logger")
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	manager.GetOrDefault("port", "9000")

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", manager.GetOrDefault("port", "9000"))))
}
