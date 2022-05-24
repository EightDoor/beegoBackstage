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
}
