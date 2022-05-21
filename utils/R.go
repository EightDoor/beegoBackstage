package utils

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type BaseController struct {
	web.Controller
}

func (c *BaseController) RSuccess(result R) {
	if result.Code == 0 {
		result.Code = 0
	}
	if result.Msg == "" {
		result.Msg = "success"
	}
	formatJsonCall(&result)
	c.Data["json"] = result
	c.ServeJSON()
}
func (c *BaseController) RError(result R) {
	if result.Code == 0 {
		result.Code = -1
	}
	if result.Msg == "" {
		result.Msg = "error"
	}
	formatJsonCall(&result)
	c.Data["json"] = result
	c.ServeJSON()
}

func formatJsonCall(result *R) {
	if _, dataErr := json.Marshal(&result.Data); dataErr != nil {
		logs.Debug(dataErr, "传入的data json 转换失败")
	}
	if _, err := json.Marshal(&result); err != nil {
		logs.Debug(err, "R json 转换失败")
	}
}
