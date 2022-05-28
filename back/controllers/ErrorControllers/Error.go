package ErrorControllers

import (
	"beegoBackstage/utils"
)

// ErrorController 定义自定义控制器
type ErrorController struct {
	utils.BaseController
}

// Error500 定义500错误，调用例子: c.Abort("500")
func (c *ErrorController) Error500() {
	c.RError(utils.R{
		Msg: "500",
	})
}

// Error404 定义404错误，调用例子: c.Abort("404")
func (c *ErrorController) Error404() {
	c.RError(utils.R{
		Msg: "404",
	})
}

// Error503 定义503错误，调用例子: c.Abort("503")
func (c *ErrorController) Error503() {
	c.RError(utils.R{
		Msg: "503",
	})
}

// Error401 定义401错误，调用例子: c.Abort("401")
func (c *ErrorController) Error401() {
	c.RError(utils.R{
		Msg: "401",
	})
}

// Error403 定义403错误，调用例子: c.Abort("403")
func (c *ErrorController) Error403() {
	c.RError(utils.R{
		Msg: "403",
	})
}
