package vo

type TodayIP struct {
	// 今日ip数量
	Count int64 `json:"count"`
}

type YesterdayIP struct {
	// 昨日ip数量
	Count int64 `json:"count"`
}

type ThisMonthIP struct {
	// 本月ip数量
	Count int64 `json:"count"`
}

type LastMonthIP struct {
	// 上月ip数量
	Count int64 `json:"count"`
}

type IPVisit struct {
	Bucket string `json:"bucket"`
	Count  int64  `json:"count"`
}

type UVVisit struct {
	Bucket string `json:"bucket"`
	Count  int64  `json:"count"`
}

type HomeIPAndUVisit struct {
	// ip 时间-数量
	IP []IPVisit `json:"ip"`
	// uv 时间-数量
	UV []UVVisit `json:"uv"`
}

type RegionStatisticResp struct {
	// 省份
	Region string `json:"region"`
	// ip个数
	Count int64 `json:"count"`
}

type JSVisitStatisticResp struct {
	Title string `json:"title"`
	Rank  int    `json:"rank"`
	Count int64  `json:"count"`
}
