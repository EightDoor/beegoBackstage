package middleware

import (
	"beegoBackstage/models"
	"beegoBackstage/utils"
	"github.com/beego/beego/v2/server/web/context"
)

var FilterUser = func(ctx *context.Context) {
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
