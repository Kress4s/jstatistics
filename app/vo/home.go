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
	// 访问量
	Count int64 `json:"count"`
	// 时间
	Bucket string `json:"bucket"`
}

type UVVisit struct {
	// 访问量
	Count int64 `json:"count"`
	// 时间
	Bucket string `json:"bucket"`
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
	// 排名
	Rank int `json:"rank"`
	// 标题
	Title string `json:"title"`
	// 访问数
	Count int64 `json:"count"`
}
