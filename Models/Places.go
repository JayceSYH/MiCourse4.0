package Models

import "github.com/astaxie/beego/orm"

type Room struct {
	Id uint64		`orm:"pk;auto"`		//房间编号
	AreaName string					//区域名 eg:仙一、逸夫楼
	RoomName string					//房间名 eg:110
	Courses []*Course	`orm:"reverse(many)"`
}

type Campus struct {
	Id uint64		`orm:"pk;auto"`		//~~
	Name string					//校区名称
	Courses []*Course	`orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Room))
	orm.RegisterModel(new(Campus))
}
