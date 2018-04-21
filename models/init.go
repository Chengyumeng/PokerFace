package models

import (
	"github.com/Chengyumeng/PokerFace/global"
)

var (
	PersonModel *personModel
)

func InitModels() {
	PersonModel = &personModel{
		global.Orm,
	}
}
