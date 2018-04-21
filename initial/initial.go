package initial

import "github.com/Chengyumeng/PokerFace/routes"

func Run(config string) {
	InitConfig(config)
	InitDB()
	InitEcho()
	routes.InitRoutes()
}
