package model

// 分页
type BasePage struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}
