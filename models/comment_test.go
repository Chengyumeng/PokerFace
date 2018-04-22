package models_test

import (
	"testing"

	"github.com/Chengyumeng/PokerFace/config"
	"github.com/Chengyumeng/PokerFace/global"
	"github.com/Chengyumeng/PokerFace/initial"
	. "github.com/Chengyumeng/PokerFace/models"
)

func TestInsertComment(t *testing.T) {
	global.Configuration = config.NewDefaultConfiguration()
	initial.InitDB()
	InitModels()

	comment, err := CommentModel.Insert(&Comment{
		Innominate: true,
		Signature:  "孙悟空",
		UserId:     1024,
		PersonId:   16,
		Viewpoint:  "还不错！",
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(comment)
}
