package utils

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type rPaging struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data *rPagingData `json:"data"`
}

type rPagingData struct {
	Count    int         `json:"count"`
	PageSize int         `json:"pageSize"`
	PageNum  int         `json:"pageNum"`
	list     interface{} `json:"data"`
}

type BaseController struct {
	web.Controller
}

// RSuccess 成功返回
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

// RError 失败返回
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

// RPaging 分页返回
func (c *BaseController) RPaging(o orm.QuerySeter) {
	result := new(rPaging)
	rData := new(rPagingData)
	page := 10
	size := 20
	count := 0
	//// 页数
	//pageNum, pageNumError := c.GetInt("pageNum")
	//if pageNumError == nil {
	//	page = pageNum
	//}
	//// 每页条数
	//pageSize, pageSizeError := c.GetInt("pageSize")
	//if pageSizeError == nil {
	//	size = pageSize
	//}
	//result.Code = 0
	//result.Msg = "success"
	//qs := o.Limit(page).Offset(size)
	var dataInfo []interface{}
	o.All(&dataInfo)
	//callCount, callCountError := qs.Count()
	//if callCountError == nil {
	//	count = int(callCount)
	//}
	rData.list = dataInfo
	rData.PageNum = page
	rData.PageSize = size
	rData.Count = count
	result.Data = rData
	formatJsonRPagingCall(result)
	c.Data["json"] = result
	c.ServeJSON()
}

// 格式化json R
func formatJsonCall(result *R) {
	if _, dataErr := json.Marshal(&result.Data); dataErr != nil {
		logs.Debug(dataErr, "传入的data json 转换失败")
	}
	if _, err := json.Marshal(&result); err != nil {
		logs.Debug(err, "R json 转换失败")
	}
}

// 格式化json R
func formatJsonRPagingCall(result *rPaging) {
	if _, dataErr := json.Marshal(&result.Data); dataErr != nil {
		logs.Debug(dataErr, "传入的data json 转换失败")
	}
	if _, err := json.Marshal(&result); err != nil {
		logs.Debug(err, "R json 转换失败")
	}
}
