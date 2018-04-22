package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type personModel struct {
	orm *xorm.Engine
}

// 评测人物
type Person struct {
	Id         int64     `xorm:"int pk autoincr"`
	Name       string    `xorm:"varchar(1024)" json:"name"`
	RecordId   int64     `xorm:"int"`
	University string    `xorm:varchar(1024) json:"university"`
	Academy    string    `xorm:varchar(1024) json:"academy"`
	Laboratory string    `xorm:varchar(1024) json:"laboratory"`
	Title      string    `xorm:"varchar(1024)" json:"title"`
	Intro      string    `xorm:"text" json:"intro"`
	CreatedAt  time.Time `xorm:"created" json:"createdAt,omitempty"`
	UpdateAt   time.Time `xorm:"created" json:"updateAt,omitempty"`
}

func (m *personModel) Insert(person *Person) (*Person, error) {
	_, err := m.orm.InsertOne(person)
	if err != nil {
		return nil, err
	}
	return person, nil
}
