package global

import (
	"github.com/labstack/echo"

	"github.com/Chengyumeng/PokerFace/config"
	"github.com/go-xorm/xorm"
)

var (
	// Echo 应用对象
	Echo *echo.Echo
	// 配置文件对象 toml描述
	Configuration *config.Configuration
	// 主数据库的对象关系映射
	Orm *xorm.Engine
)
