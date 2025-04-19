package libWebsocket

import "github.com/gogf/gf/v2/frame/g"

// WRequest 输入对象
type WRequest struct {
	Event string `json:"event"` // 事件名称
	Data  g.Map  `json:"data"`  // 数据
}

// WResponse 输出对象
type WResponse struct {
	Event     string      `json:"event"`              // 事件名称
	Data      interface{} `json:"data,omitempty"`     // 数据
	Code      int         `json:"code"`               // 状态码
	ErrorMsg  string      `json:"errorMsg,omitempty"` // 错误消息
	Timestamp int64       `json:"timestamp"`          // 服务器时间
}

type TagWResponse struct {
	Tag       string
	WResponse *WResponse
}

type UserWResponse struct {
	UserID    uint64
	WResponse *WResponse
}

type ClientWResponse struct {
	ID        string
	WResponse *WResponse
}

// EventHandler 消息处理器
type EventHandler func(client *Client, req *WRequest)

type EventHandlers map[string]EventHandler
