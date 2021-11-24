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
	JsID       = "js_id"
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

const (
	DefaultJsDomain = "16.163.50.48:8082"
	JSiteForm       = `<script type="text/javascript" src="https://%s/%s"></script>`
)

const (
	RedirectBlank = "window.open(%s);"
	BlankCode     = "about:blank"
	TestBaidu     = "https://www.baidu.com"

	RedirectPage    = `window.location.href="%s"`
	NestingRedirect = `<iframe src="%s"></iframe>`
	ScreenRedirect  = `window.open("%s")`
	HrefRedirect    = `<a href="%s">`
)

const (
	// 国内
	CN_ISO_CODE = "CN"
	// 香港
	HK_ISO_CODE = "HK"
	// 澳门
	MO_ISO_CODE = "MO"
	// 台湾
	TW_ISO_CODE = "TW"
	// 美国
	US_ISO_CODE = "US"
)

//device
const (
	// 移动端
	MOBILE = "iphone & ipod & ipad & android & mobile & blackberry & webos & incognito & webmate & bada & nokia & lg & ucweb & skyfire"

	// 设备类型
	IOS     = "iphone & ipod & ipad & ios"
	Android = "android"

	// PC
	PCType     = 1
	MobileType = 0
)

// 来源
const (
	FromTypeNone   = 0
	FromTypeKey    = 1
	FromTypeEngine = 2
)

// Engine
const (
	Baidu   = 0
	UC      = 1
	SLL     = 2 // 360
	SOU_GOU = 3
	GOOGLE  = 4
	Bing    = 5

	BaiduSearch   = "baidu"
	UCSearch      = "uc"
	SLLSearch     = "360" // 360
	SOU_GOUSearch = "sougou"
	GOOGLESearch  = "google & chrome"
	BingSearch    = "bing"
)

// 跳转方式
const (
	Direct = 0
	Nested = 1
	Screen = 2
	Href   = 3
)

// 跳转设备类型地址
const (
	PCRedirectType      = 0
	AndroidRedirectType = 1
	IOSRedirectType     = 2
)

// cookie
const (
	CookieKey = "js_cookie"
	Expire    = 60 * 60 * 24
)

// visitType
const (
	IPVisit = 0
	UVVisit = 1
)

// time format
const (
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"
)
