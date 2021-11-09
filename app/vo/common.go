package vo

// Pagination 分页信息
type Pagination struct {
	// 请求页
	Page int `json:"page"`
	// 页大小
	PageSize int `json:"page_size"`
	// 数据总条数
	TotalCount int64 `json:"total_count"`
}

// DataPagination 数据包含分页信息
type DataPagination struct {
	// 数据
	Data interface{} `json:"data"`
	// 分页信息
	Pagination Pagination `json:"pagination"`
}

type PageInfo struct {
	// 请求页
	Page int `json:"page"`
	// 页大小
	PageSize int `json:"page_size"`
	// 模糊搜索关键字
	Keywords string `json:"keywords"`
}

func (p *PageInfo) Offset() int {
	return (p.Page - 1) * p.PageSize
}

func NewDataPagination(count int64, data interface{}, page *PageInfo) *DataPagination {
	return &DataPagination{
		Data: data,
		Pagination: Pagination{
			Page:       page.Page,
			PageSize:   page.PageSize,
			TotalCount: count,
		},
	}
}
