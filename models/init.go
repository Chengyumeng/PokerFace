package models

import (
	"github.com/Chengyumeng/PokerFace/global"
)

var (
	PersonModel *personModel
	UserModel   *userModel
)

func InitModels() {
	PersonModel = &personModel{
		global.Orm,
	}
	UserModel = &userModel{
		global.Orm,
	}

}
