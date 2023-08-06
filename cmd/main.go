package main

import (
	"github.com/Edwinfpirajan/practice.git/config"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectionDatabase()

	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
}
