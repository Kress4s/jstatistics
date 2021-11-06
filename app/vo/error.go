package vo

type Error struct {
	// 错误码
	Code int `json:"code"`
	// 错误消息
	Msg string `json:"msg"`
	// 参数
	Args []string `json:"args"`
}
