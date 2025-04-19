package consts

import "github.com/gogf/gf/v2/errors/gcode"

var (
	ServerError = gcode.New(1001, "服务异常，请稍后重试", nil)
)
