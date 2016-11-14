package Models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:QQ140100210031s@tcp(115.159.223.65:3306)/miktest?charset=utf8", 30)
}
