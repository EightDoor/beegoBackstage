package BaseControllers

import beego "github.com/beego/beego/v2/server/web"

type LoginController struct {
	beego.Controller
}

// Get
// @router / [get]
func (c *LoginController) Get() {
	c.Data["json"] = map[string]string{"login": "true"}
	c.ServeJSONP()
}
