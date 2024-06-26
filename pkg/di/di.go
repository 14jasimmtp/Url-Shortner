package di

import (
	"github.com/14jasimmtp/Url-Shortner/pkg/controller"
	"github.com/14jasimmtp/Url-Shortner/pkg/db"
	"github.com/14jasimmtp/Url-Shortner/pkg/routes"
	"github.com/labstack/echo/v4"
)

func InitialiseAPI(e *echo.Echo) {
	DB:=db.ConnectToDB()
	controller:=controller.NewController(DB)
	routes.Routes(e,controller)
}