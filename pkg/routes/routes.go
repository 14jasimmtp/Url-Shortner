package routes

import (
	"github.com/14jasimmtp/Url-Shortner/pkg/controller"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo,c *controller.UrlHandler) {
	e.GET("/:ShotKey",c.GetUrl)
	e.POST("/shortURL",c.ShortUrl)
}