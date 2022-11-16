package main

import (
	"fmt"
	"github.com/go-nag/configuration/cfgl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

var cfg, _ = cfgl.LoadConfigFile("local")

func main() {
	e := echo.New()

	loggerEnabled, err := cfg.Get("server.logging")
	if err != nil {
		log.Fatal(err)
	}

	if loggerEnabled == "enabled" {
		log.Println("Using logger")
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.GetOrDefault("port", "9000"))))
}
