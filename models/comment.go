package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type commentModel struct {
	orm *xorm.Engine
}

// 评论
type Comment struct {
	Id         int64  `xorm:"int pk autoincr"`
	PersonId   int64  `xorm:"int"`
	Innominate bool   `xorm:"bool" json:"innominate"`         // 是否匿名，默认匿名
	Signature  string `xorm:"varchar(1024)" json:"signature"` // 仅代表签名，只有非匿名状态才会显示真实用户ID
	UserId     int64  `xorm："int"`                            // 登录的匿名用户将被记录ID，如非选择，不会公开真实信息
	Viewpoint  string `xorm："text" json:"viewpoint"`          // 评价内容
	VoteUp     int64  `xorm:"int" json:"vote_up"`             // 认同次数
	VoteDown   int64  `xorm:"int" json:"vote_down"`           // 否决次数
	Score      int64  `xorm:"int" json:"score"`               // 评分

	CreatedAt time.Time `xorm:"created" json:"createdAt,omitempty"`
	UpdateAt  time.Time `xorm:"created" json:"updateAt,omitempty"`
}

func (m *commentModel) Insert(comment *Comment) (*Comment, error) {
	_, err := m.orm.InsertOne(comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
