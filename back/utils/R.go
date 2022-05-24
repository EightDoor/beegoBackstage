package utils

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type RPage struct {
	Code      int            `json:"code"`
	Msg       string         `json:"msg"`
	Data      interface{}    `json:"data"`
	OrmResult orm.QuerySeter `json:"ormResult"`
}

type rPaging struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data *rPagingData `json:"data"`
}

type rPagingData struct {
	Total    int         `json:"total"`
	PageSize int         `json:"pageSize"`
	PageNum  int         `json:"pageNum"`
	List     interface{} `json:"list"`
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

// RBack 返回内容
func (c *BaseController) RBack(result R, err error) {
	if err != nil {
		result.Msg = err.Error()
		c.RError(result)
	} else {
		c.RSuccess(result)
	}
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
// TODO 待改进，是否可以去掉dataInfo参数，去掉存在问题(根据反射获取不到对应的struct结构)
func (c *BaseController) RPaging(rPage RPage) {
	result := new(rPaging)
	rData := new(rPagingData)
	page := 1
	size := 20
	// 读取分页默认配置
	defaultPageNum, defaultPageNumErr := web.AppConfig.String("pageNum")
	if defaultPageNumErr == nil {
		page, _ = strconv.Atoi(defaultPageNum)
	}
	defaultPageSize, defaultPageSizeErr := web.AppConfig.String("pageSize")
	if defaultPageSizeErr == nil {
		size, _ = strconv.Atoi(defaultPageSize)
	}
	count := 0
	// 页数
	pageNum, pageNumError := c.GetInt("pageNum")
	if pageNumError == nil {
		page = pageNum
	}
	// 每页条数
	pageSize, pageSizeError := c.GetInt("pageSize")
	if pageSizeError == nil {
		size = pageSize
	}
	result.Code = 0
	result.Msg = "success"

	rPage.OrmResult.Limit(size).Offset((page - 1) * size).All(rPage.Data)
	callCount, callCountError := rPage.OrmResult.Count()
	if callCountError == nil {
		count = int(callCount)
	}
	rData.List = rPage.Data
	rData.PageNum = page
	rData.PageSize = size
	rData.Total = count
	result.Data = rData
	formatJsonRPagingCall(result)
	c.Data["json"] = result
	c.ServeJSON()
}

// 格式化json R
func formatJsonCall(result *R) {
	if _, dataErr := json.Marshal(&result.Data); dataErr != nil {
		logs.Error(dataErr, "传入的data json 转换失败")
	}
	if _, err := json.Marshal(&result); err != nil {
		logs.Error(err, "R json 转换失败")
	}
}

// 格式化json rPaging
func formatJsonRPagingCall(result *rPaging) {
	if _, dataErr := json.Marshal(&result.Data); dataErr != nil {
		logs.Error(dataErr, "传入的data json 转换失败")
	}
	if _, err := json.Marshal(&result); err != nil {
		logs.Error(err, "R json 转换失败")
	}
}

// GetBodyToJson
// body接受参数解析 JSON
func (c *BaseController) GetBodyToJson(data interface{}) {
	err := json.Unmarshal(c.Ctx.Input.RequestBody, data)
	if err != nil {
		c.RError(R{Msg: err.Error()})
	}
}

// CustomValid 表单校验
func (c *BaseController) CustomValid(data interface{}) {
	valid := validation.Validation{}
	rValid, rValidErr := valid.Valid(data)
	if rValidErr != nil {
		logs.Error(rValidErr, data, "表单校验失败")
		c.RError(R{
			Msg: rValidErr.Error(),
		})
	}

	if !rValid {
		for _, err := range valid.Errors {
			logs.Info(err.Key, err.Message)
		}
		c.RError(R{
			Data: valid.Errors,
		})
	}
}
