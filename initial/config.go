package initial

import (
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"

	"github.com/Chengyumeng/PokerFace/config"
	"github.com/Chengyumeng/PokerFace/global"
)

func InitConfig(filename string) {
	configAbsPath, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(configAbsPath)
	if err != nil {
		panic(err)
	}

	conf := config.NewDefaultConfiguration()

	if _, err := toml.Decode(string(data), conf); err != nil {
		panic(err)
	}

	global.Configuration = conf
}
