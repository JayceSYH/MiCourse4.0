package Models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Message struct {
	Id uint64		`orm:"pk;auto"`		//~~
	Addressor string				//发送者
	Addressee string 				//收件人
	IsRead uint8					//是否已阅
	AddTime time.Time				//创建时间
}


type Notification struct {
	Id uint64		`orm:"pk;auto"`		//~~
	User *User		`orm:"rel(fk)"`
	NotificationType string				//通知类型
	Content string 					//通知内容
	AddTime time.Time				//创建时间
	IsRead uint8					//是否已阅
}


func init() {
	orm.RegisterModel(new(Message))
	orm.RegisterModel(new(Notification))
}