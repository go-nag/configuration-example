package main

import (
	"fmt"
	"github.com/go-nag/configuration/cfgm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

var cfg, _ = cfgm.LoadConfigFile("local")

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

	arr := cfg.GetArr("array.value")
	// Logs 2022/12/10 16:52:55 [test1 test2 test3]
	log.Println(arr)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.GetOrDefault("port", "9000"))))
}
