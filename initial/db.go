package initial

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"github.com/Chengyumeng/PokerFace/global"
)

func InitDB() {
	database := global.Configuration.AppConfig.Database

	dbUrl := fmt.Sprintf(
		"%s:%s@%s?charset=%s&parseTime=true",
		database.User,
		database.Password,
		database.Tns,
		database.Charset,
	)
	orm, err := xorm.NewEngine(database.DriverName, dbUrl)
	if err != nil {
		panic(err)
	}

	orm.SetMaxIdleConns(0)
	orm.ShowSQL(database.ShowSQL)

	global.Orm = orm
}
