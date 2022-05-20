package controllers

import beego "github.com/beego/beego/v2/server/web"

type LoginController struct {
	beego.Controller
}

// GET @router / [get]
func (c *LoginController) GET() {
	c.Data["json"] = map[string]string{"login": "true"}
	c.ServeJSONP()
}
