package vo

type FlowDataResp struct {
	// 标题
	Title string `json:"title"`
	// ip总数
	IP int64 `json:"ip"`
	// uv总数
	UV int64 `json:"uv"`
}
