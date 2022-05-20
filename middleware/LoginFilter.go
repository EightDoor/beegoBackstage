package middleware

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
)

var FilterUser = func(ctx *context.Context) {
	//_, ok := ctx.in
	logs.Info("进入登录校验!")
}
