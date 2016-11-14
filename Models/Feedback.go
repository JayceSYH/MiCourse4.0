package Models

import "time"

type Feedback struct {
	Id uint64		`orm:"pk;auto"`		//~~
	Email string					//邮箱
	Content string					//内容
	AddTime time.Time				//创建时间
}
