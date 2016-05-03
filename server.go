package maple

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/tylerb/graceful"

	"github.com/jaem/maple/routers"
	"github.com/jaem/nimble"
)

func InitDatabase(c *config) {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(c.Db_Alias, "mysql", c.Db_Auth, c.Db_MaxIdle, c.Db_MaxConn)
	err := orm.RunSyncdb(c.Db_Alias, false, c.DebugMode)
	if err != nil {
		fmt.Println(err)
	}
}

func StartServer(c *config) {
	web := nimble.Default()
	web.Use(routers.InitRoutes())

	if c.DebugMode {
		web.Run(c.Server_Port) // Development
	} else {
		graceful.Run(c.Server_Port, c.Server_Timeout, web) // Production
	}
}
