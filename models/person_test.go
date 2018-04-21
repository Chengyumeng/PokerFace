package models_test

import (
	"testing"

	"github.com/Chengyumeng/PokerFace/config"
	"github.com/Chengyumeng/PokerFace/global"
	"github.com/Chengyumeng/PokerFace/initial"
	. "github.com/Chengyumeng/PokerFace/models"
)

func TestInsertPerson(t *testing.T) {
	global.Configuration = config.NewDefaultConfiguration()
	initial.InitDB()
	InitModels()

	person, err := PersonModel.Insert(&Person{
		Name:       "Chengyumeng",
		University: "清华大学",
		RecordId:   1,
		Academy:    "软件学院",
		Laboratory: "人工智能实验室",
		Title:      "研究员",
		Intro:      "上面的信息都是假的。",
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(person)
}
