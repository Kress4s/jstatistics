package vo

type FromAnalysisResp struct {
	Title string `json:"title"`
	From  string `json:"from"`
	To    string `json:"to"`
	Count int64  `json:"count"`
}
