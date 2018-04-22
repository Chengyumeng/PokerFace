package initial

import (
	"github.com/Chengyumeng/PokerFace/models"
	"github.com/Chengyumeng/PokerFace/routes"
)

func Run(config string) {
	InitConfig(config)
	InitDB()
	models.InitModels()
	InitEcho()
	routes.InitRoutes()
}
