package global

import (
	"github.com/labstack/echo"

	"github.com/Chengyumeng/PokerFace/config"
)

var (
	// Echo 应用对象
	Echo *echo.Echo
	// 配置文件对象 toml描述
	Configuration *config.Configuration
)
