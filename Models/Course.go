package Models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

//课程信息
type Course struct {
	Id uint64 		   `orm:"pk;auto"`			//米课内部课程编号
	CourseCode uint64	   					//官方课程编码
	CourseName string						//课程名
	Credit uint8							//学分
	Teacher string							//教师名字
	Category string							//课程类型
	StartWeek uint8							//开课周
	EndWeek uint8							//最后一周
	State uint8							//课程信息信息状态(课程关闭、失效 or sth)
	AddTime time.Time						//创建时间
	LastTime time.Time						//上次更新时间
	CourseDesc *CourseDesc	   `orm:"reverse(one)"`		//课程描述
	CourseScore []*CourseScore `orm:"reverse(many)"`		//课程评分
	CourseAssistant *CourseAssistant `orm:"reverse(one)"`		//课程助教
	CourseStatistic []*CourseStatistic `orm:"reverse(many)"`	//课程统计数据
	CourseDepartment *CourseDepartment `orm:"rel(fk)"`		//开设课程的部门
	CourseType *CourseType   `orm:"rel(fk)"`			//课程类型编号（通识课等）
	CourseTime *CourseTime	   `orm:"reverse(one)"`			//课程时间
	Campus *Campus		   `orm:"rel(fk)"`			//开课校区
	Room *Room		   `orm:"rel(fk)"`			//房间编号
	Semester *Semester	   `orm:"rel(fk)"`			//学期编号
	Keyword []*Keyword	   `orm:"rel(m2m)"`			//关键字
	Users []*UserCourse	   `orm:"reverse(many)"`
}

//课程描述
type CourseDesc struct {
	Id uint64 		   `orm:"pk;auto"`		//米课内部课程编号
	SimpleDesc string					//课程描述
	Summery	string						//一句话简介
	Course *Course 	  	   `orm:"rel(one)"`		//对应课程
}

//课程评分
type CourseScore struct {
	Id uint64 		   `orm:"pk;auto"`		//课程评分编号
	Score uint32						//评分
	Times uint32						//课程评分次数
	Course *Course		   `orm:"rel(fk)"`
}

//课程助教
type CourseAssistant struct {
	Id uint64 		   `orm:"pk;auto"`		//助教编号
	Name string						//助教名
	Email string 						//助教邮箱
	Course *Course		   `orm:"rel(one)"`
}

//课程统计数据
type CourseStatistic struct {
	Id uint64 		   `orm:"pk;auto"`		//课程统计编号
	Topic string 						//统计主题
	Value int32						//用户设置的值
	User *User		   `orm:"rel(fk)"`
	Course *Course		   `orm:"rel(fk)"`
}

//课程时间
type CourseTime struct {
	Id uint64 		   `orm:"pk;auto"`		//~~
	Date uint8		   				//星期几上课
	StartTime uint8		   				//这天第几节开始上课
	EndTime uint8						//第几节课结束
	Course *Course		   `orm:"rel(one)"`
}

//开课部门
type CourseDepartment struct {
	Id uint64 		   `orm:"pk;auto"`			//开课部门编号
	Name string 							//部门名称
	CourseDepartmentType *CourseDepartmentType `orm:"rel(fk)"`	//部门类型
	Courses []*Course	   `orm:"reverse(many)"`
}

//开课部门类型
type CourseDepartmentType struct {
	Id uint64 		   `orm:"pk;auto"`		//~~
	Name string
	CourseDepartments []*CourseDepartment	`orm:"reverse(many)"` //
}

//课程类型
type CourseType struct {
	Id uint64 		   `orm:"pk;auto"`		//课程类型编号
	Name string						//类型名
	Courses []*Course	   `orm:"reverse(many)"`
}


func init() {
	orm.RegisterModel(new(Course))
	orm.RegisterModel(new(CourseDepartment))
	orm.RegisterModel(new(CourseStatistic))
	orm.RegisterModel(new(CourseTime))
	orm.RegisterModel(new(CourseAssistant))
	orm.RegisterModel(new(CourseScore))
	orm.RegisterModel(new(CourseDepartmentType))
	orm.RegisterModel(new(CourseDesc))
	orm.RegisterModel(new(CourseType))
}