package Models

import "github.com/astaxie/beego/orm"

type Keyword struct {
	Id uint64		`orm:"pk;auto"`			//关键字编号
	Category string						//关键字分类
	Key string						//关键字
	Course []*Course	`orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Keyword))
}