package utils

import (
	"beegoBackstage/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
)

func init() {
	// 拦截错误
	beego.BConfig.RecoverFunc = func(context *context.Context, config *beego.Config) {
		if err := recover(); err != nil {
			context.Output.Status = http.StatusInternalServerError
			context.JSONResp(R{Msg: fmt.Sprintf("发生错误, 错误信息: %v", err), Code: models.UNKNOWN_ERROR})
		}
	}
}
