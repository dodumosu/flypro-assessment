package dto

type PaginationParams struct {
	Page int `query:"page"`
}

type PageInfo struct {
	CurrentPage int `json:"currentPage"`
	Total       int `json:"total"`
	FirstPage   int `json:"firstPage"`
	LastPage    int `json:"lastPage"`
}
