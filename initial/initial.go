package initial

import (
	"github.com/Chengyumeng/PokerFace/routes"
	"github.com/Chengyumeng/PokerFace/models"
)

func Run(config string) {
	InitConfig(config)
	InitDB()
	models.InitModels()
	InitEcho()
	routes.InitRoutes()
}
