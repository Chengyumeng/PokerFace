package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type userModel struct {
	orm *xorm.Engine
}

// 网站用户
type User struct {
	Id         int64     `xorm:"int pk autoincr"`
	Name       string    `xorm:"varchar(1024)" json:"name"`
	Email      string    `xorm:"varchar(1024)" json:"email"`
	Phone      string    `xorm:"varchar(1024)" json:"phone"`
	University string    `xorm:"varchar(1024)" json:"university"`
	Grade      int32     `xorm:"int"`
	IsValid    bool      `xorm:"bool"`
	CreatedAt  time.Time `xorm:"created" json:"createdAt,omitempty"`
	UpdateAt   time.Time `xorm:"created" json:"updateAt,omitempty"`
}

func (m *userModel) Insert(user *User) (*User, error) {
	_, err := m.orm.InsertOne(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
