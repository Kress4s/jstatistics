package vo

type FromAnalysisResp struct {
	// 标题
	Title string `json:"title"`
	// 来路URL
	From string `json:"from"`
	// 去路URL
	To string `json:"to"`
	// 次数
	Count int64 `json:"count"`
}
