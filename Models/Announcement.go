package Models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

//公告
type Announcement struct {
	Id uint64		`orm:"pk;auto"`		//~~
	AnnouncementType string				//公告类型
	Title string					//标题
	Abstract string					//内容摘要
	Content string					//公告内容
	AddTime	time.Time				//创建时间
}

func init() {
	orm.RegisterModel(new(Announcement))
}
