package Models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户
type User struct {
	Id uint64					`orm:"pk;auto"`			//用户Id
	Role string 									//用户角色
	UserName string 								//用户名
	Gender uint8									//性别
	Grade string									//年级
	SchoolDepartment *SchoolDepartment		`orm:"rel(fk)"`			//院系
	Password string 								//密码
	Email string 									//邮箱
	PortraitUrl string								//用户头像
	SelfDesc string 								//自我描述
	IsAgreePush string								//是否接受推送
	IsValidated uint8								//是否已经过邮箱验证
	AddTime time.Time				`orm:"auto_now"`		//注册时间
	AddIp string									//注册IP
	CourseStatistic []*CourseStatistic		`orm:"reverse(many)"`
	Courses []*UserCourse				`orm:"rel(m2m);"`
	Friends []*UserFriend				`orm:"reverse(many)"`
	SentApplies []*UserFriendApply			`orm:"reverse(many)"`		//用户已发送的申请
	ReceivedApplies []*UserFriendApply		`orm:"reverse(many)"`
}

//用户课程关系
type UserCourse struct {
	Id uint64					`orm:"pk;auto"`			//~~
	User *User					`orm:"rel(fk)"`
	Course *Course					`orm:"rel(fk)"`
	AgreePush uint8									//是否允许推送
	Status uint8									//是否上过，是否想上
	Delete uint8									//是否被删除
	AddTime time.Time								//创建时间
}

//用户好友
type UserFriend struct {
	Id uint64					`orm:"pk;auto"`			//~~
	User *User					`orm:"rel(fk)"`
	Friend *User					`orm:"rel(fk)"`
	AddTime time.Time
}

//好友申请
type UserFriendApply struct {
	Id uint64					`orm:"pk;auto"`			//~~
	User *User					`orm:"rel(fk)"`
	Friend *User					`orm:"rel(fk)"`
	IsPassed uint8									//是否通过
	AddTime time.Time
}

//找回账号
type UserRetrieve struct {
	Id uint64					`orm:"pk;auto"`			//~~
	User *User					`orm:"rel(fk)"`
	RetrieveKey string								//校验码
	Success uint8									//是否成功
	AddTime time.Time								//创建时间
}


//用户验证
type UserValidation struct {
	Id uint64					`orm:"pk;auto"`			//~~
	User *User					`orm:"rel(fk)"`
	ValidationKey string 								//验证秘钥
	Expired time.Time								//过期时间
}


//换课信息
type ExchangeCourse struct {
	Id uint64					`orm:"pk;auto"`			//~~
	User *User					`orm:"rel(fk)"`
	Course *Course					`orm:"rel(fk)"`
	Title string									//标题
	Desc string									//一段描述
	State uint8									//交换状态
}


func init() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(UserCourse))
	orm.RegisterModel(new(UserFriend))
	orm.RegisterModel(new(UserFriendApply))
	orm.RegisterModel(new(UserRetrieve))
	orm.RegisterModel(new(UserValidation))
	orm.RegisterModel(new(ExchangeCourse))
}