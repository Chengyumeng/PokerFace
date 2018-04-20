package routes

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/Chengyumeng/PokerFace/global"
)

func InitRoutes() {
	global.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
