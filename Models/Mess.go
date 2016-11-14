package Models

import "github.com/astaxie/beego/orm"

type SchoolDepartment struct {
	Id uint64		`orm:"pk;auto"`			//~~
	Name string
}

func init() {
	orm.RegisterModel(new(SchoolDepartment))
}
