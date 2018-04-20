package main

import (
	"github.com/Chengyumeng/PokerFace/global"
	"github.com/Chengyumeng/PokerFace/initial"
)

func main() {
	initial.Run("hack/porkerface.toml")

	addr := global.Configuration.AppConfig.Address

	global.Echo.Logger.Fatal(global.Echo.Start(addr))

}
