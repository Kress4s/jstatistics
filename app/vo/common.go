package vo

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

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

func FuzzySearch(keywords, key string, moreKeys ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		condition := make([]string, 0)
		keywords = strings.ReplaceAll(keywords, "|", "||")
		keywords = strings.ReplaceAll(keywords, "_", "|_")
		keywords = strings.ReplaceAll(keywords, "%", "|%")
		keywords = strings.ReplaceAll(keywords, "'", "|'")
		searchText := "%" + strings.ToUpper(keywords) + "%"
		condition = append(condition, fmt.Sprintf(`upper(%s) LIKE ? ESCAPE '|'`, key))
		for _, v := range moreKeys {
			condition = append(condition, fmt.Sprintf(`upper(%s) LIKE ? ESCAPE '|'`, v))
		}
		values := make([]interface{}, len(moreKeys)+1)
		for i := range values {
			values[i] = searchText
		}
		return db.Where(strings.Join(condition, " OR "), values...)
	}
}
