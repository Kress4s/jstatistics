package vo

type ObjectResp struct {
	// 对象id
	ID string `json:"id"`
	// 对象名称
	Filename string `json:"filename"`
	// 对象内容
	Content string `json:"content"`
}

type UUID struct {
	ID string `json:"id"`
}
