package Models

import "github.com/astaxie/beego/orm"

type Semester struct {
	Id uint64		`orm:"pk;auto"`		//学期编号
	Name string					//学期名称
	Courses []*Course	`orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Semester))
}