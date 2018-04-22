package models_test

import (
	"testing"

	"github.com/Chengyumeng/PokerFace/config"
	"github.com/Chengyumeng/PokerFace/global"
	"github.com/Chengyumeng/PokerFace/initial"
	. "github.com/Chengyumeng/PokerFace/models"
)

func TestInsertUser(t *testing.T) {
	global.Configuration = config.NewDefaultConfiguration()
	initial.InitDB()
	InitModels()

	user, err := UserModel.Insert(&User{
		Name:       "Chengyumeng",
		University: "清华大学",
		Email:      "chengyumeng@gmail.com",
		IsValid:    true,
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}
