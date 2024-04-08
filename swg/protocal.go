package swg

type CodeMsg struct {
	Code int    `json:"code" example:"0"`      // 响应码 非0即为失败
	Msg  string `json:"msg" example:"success"` // msg
}

type Page struct {
	TotalCount int64  `json:"total_count" example:"100"` // 列表数据总数
	TotalPage  int64  `json:"total_page" example:"10"`   // 列表页总数
	PageCount  int64  `json:"page_count" example:"10"`   // 当前页数量
	PageIndex  int64  `json:"page_index" example:"3"`    // 当前页码
	Sort       string `json:"sort"`                      // 排序
}

type SingleData[T any] struct {
	CodeMsg
	Data T `json:"data"`
}

type MultiData[T any] struct {
	CodeMsg
	Data []T `json:"data"`
}
