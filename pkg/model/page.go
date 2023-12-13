package model

type Paging[T any] struct {
	PageNumber int   `json:"pageNumber"`      // 页号
	PageSize   int   `json:"pageSize"`        // 分页 默认10
	Total      int64 `json:"total,omitempty"` // 总页数
	Data       []T   `json:"data,omitempty"`  // 数据项
}

type Search struct {
	PageNumber int      `json:"pageNumber"`        // 页号
	PageSize   int      `json:"pageSize"`          // 分页 默认10
	Keyword    string   `json:"keyword,omitempty"` // 关键字
	Orders     []*Order `json:"orders,omitempty"`  // 排序
}

type Order struct {
	Asc    bool   `json:"asc"`    // 降序
	Column string `json:"column"` // 字段
}
