package Models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

//招新信息
type Recruit struct {
	Id uint64		`orm:"pk;auto"`			//~~
	Name string
	Gender string
	Grade string
	SchoolDepartment *SchoolDepartment	`orm:"rel(fk)"`	//院系
	Email string
	Mobile string
	LanguageExperience string
	DesignExperience string
	Document string
	ideas string
	AddTime time.Time
}

func init()  {
	orm.RegisterModel(new(Recruit))
}