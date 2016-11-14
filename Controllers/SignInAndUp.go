package Controllers

import (
	"github.com/astaxie/beego"
)

type SignController struct {
	beego.Controller
}

func (c *SignController) Get() {
	c.Ctx.WriteString("sign up")
}

func (c *SignController) Post() {
	c.Ctx.WriteString("sign in")
}

func init() {
	beego.Router("/user/sign", &SignController{})
}

