package initial

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/Chengyumeng/PokerFace/global"
)

func InitEcho() {
	global.Echo = echo.New()

	global.Echo.Use(middleware.Logger())
	global.Echo.Use(middleware.Recover())
}
