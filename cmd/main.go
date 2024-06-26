package main

import (
	"github.com/14jasimmtp/Url-Shortner/pkg/di"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	di.InitialiseAPI(e)
	e.Logger.Fatal(e.Start(":3000"))
}
