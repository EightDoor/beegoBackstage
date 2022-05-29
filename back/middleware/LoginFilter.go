package middleware

import (
	"beegoBackstage/models"
	"beegoBackstage/models/JournalModels"
	"beegoBackstage/models/SysModels"
	"beegoBackstage/utils"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
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
	userId, err := utils.ValidateToken(token)
	if err != nil {
		r.Msg = err.Error()
		ctx.JSONResp(r)
	}

	// TODO 正常开发可以去掉
	// 读取配置值
	env, _ := web.AppConfig.String("runmode")
	if env == "prod" {
		// TODO 过滤PUT、DELETE、POST 不能操作
		// 超级管理员不可以操作
		if userId == 11 {
			if ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE" || ctx.Request.Method == "POST" {
				r.Msg = "演示版本，不可以操作"
				r.Data = ""
				r.Code = models.DEMO_ENV
				ctx.JSONResp(r)
			}
		}
	}

	logs.Info(ctx.Request.Body, "Body")
	logs.Info(ctx.Request.MultipartForm, "ParamsForm")
	userLog, _ := web.AppConfig.Bool("userLog")
	if userLog {
		var logRequest JournalModels.LogRequest
		var sysUser SysModels.SysUser
		sysUser.Id = userId
		logRequest.Ip = ctx.Input.Header("customIp")
		logRequest.User = &sysUser
		isUrlIdentification := strings.Index(ctx.Request.RequestURI, "?")
		if isUrlIdentification == -1 {
			logRequest.RequestAddress = ctx.Request.RequestURI
		} else {
			logRequest.RequestAddress = ctx.Request.RequestURI[0:isUrlIdentification]
		}

		logRequest.Method = ctx.Request.Method
		// 解析body
		var m map[string]interface{}
		json.NewDecoder(ctx.Request.Body).Decode(&m)
		resultM, _ := json.Marshal(m)
		resultMData := string(resultM)
		logs.Info(resultM, "body - resultM")
		if logRequest.Method != "POST" {
			ctx.Request.ParseForm()
			values := ctx.Request.Form
			valuesData, _ := json.Marshal(values)
			logRequest.Params = string(valuesData)
		} else {
			logRequest.Params = resultMData
		}
		JournalModels.LogRequestCreate(logRequest)
	}
}
