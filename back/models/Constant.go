package models

// 状态码 100开始
const (
	// 未知错误
	UNKNOWN_ERROR = (iota + 100)
	// token验证失败 101
	NO_AUTHORIZATION
	// 演示环境
	DEMO_ENV
)
