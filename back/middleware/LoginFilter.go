package middleware

import (
	"beegoBackstage/models"
	"beegoBackstage/utils"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"strings"
)

var FilterUser = func(ctx *context.Context) {
	whiteList := [...]string{"/login"}
	url := ctx.Request.RequestURI
	for _, v := range whiteList {
		if strings.Index(url, v) != -1 {
			return
		}
	}
	logs.Info(url, "请求url")
	var r utils.R
	token := ctx.Input.Header("Authorization")
	r.Code = models.NO_AUTHORIZATION
	// 验证token
	_, err := utils.ValidateToken(token)
	if err != nil {
		r.Msg = err.Error()
		ctx.JSONResp(r)
	}

	// TODO 过滤PUT、DELETE、POST 不能操作
	if ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE" || ctx.Request.Method == "POST" {
		r.Msg = "演示版本，不可以操作"
		r.Data = ""
		r.Code = models.DEMO_ENV
		ctx.JSONResp(r)
	}
}
