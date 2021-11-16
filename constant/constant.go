package constant

// Auth
const (
	Salt          = "JS Secret"
	Authorization = "Bearer"
)

// tableName
const (
	CostCategory = "category"
	Account      = "account"
)

// http request
const (
	ID         = "id"
	IDS        = "ids"
	PrimaryID  = "pid"
	CategoryID = "cid"
	RoleID     = "role_id"
	BeginAt    = "begin_at"
	EndAt      = "end_at"
	Year       = "year"
	Date       = "date"
	TimeAt     = "time_at"
)

// pagination key
const (
	Page       = "page"
	PageSize   = "page_size"
	TextSearch = "keywords"
)

const (
	// jsp 文件ip处理，为解决
	// IPLocation = "http://whois.pconline.com.cn/ipJson.jsp?ip=%s&json=true"

	IPLocation = "http://ip-api.com/json/%s?lang=zh-CN"
	IP         = "ip"
)
