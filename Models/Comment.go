package Models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type CourseComment struct {
	Id uint64		`orm:"pk;auto"`			//id
	User *User		`orm:"rel(fk)"`
	Course *Course		`orm:"rel(fk)"`
	State uint8						//评论状态
	Useful uint8						//评论有用
	Useless uint8						//评论无用
	Content string						//评论内容
	AddTime	time.Time					//创建时间
}


type CommentReply struct {
	Id uint64			`orm:"pk;auto"`		//回复编号
	CourseComment *CourseComment	`orm:"rel(one)"`	//原评论编号
	User *User			`orm:"rel(fk)"`		//回复者
	ParentComment *CommentReply	`orm:"rel(fk)"`		//若是回复的回复，编号为父回复的编号
	SubComment []*CommentReply	`orm:"reverse(many)"`	//子评论
	State uint8						//评论状态
	Content string						//评论内容
	AddTime time.Time					//评论时间
}


type UserVote struct {
	Id uint64			`orm:"pk;auto"`		//~~
	User *User			`orm:"rel(fk)"`		//投票用户
	CourseComment *CourseComment	`orm:"rel(fk)"`
	IsSupported uint8
	AddTime time.Time					//创建时间
}


func init() {
	orm.RegisterModel(new(CourseComment))
	orm.RegisterModel(new(CommentReply))
	orm.RegisterModel(new(UserVote))
}