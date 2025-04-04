package presenters

type Pagination struct {
	PageSize      int `json:"page_size"`
	Page          int `json:"page"`
	TotalElements int `json:"total_elements"`
}

type ResponseNoPagination struct {
	Results any `json:"results"`
}

type ResponsePagination struct {
	Pagination Pagination `json:"pagination"`
	Results    any        `json:"results"`
}

func BuildResponsePagination(page, pageSize, totalElements int, data any) ResponsePagination {
	return ResponsePagination{
		Pagination: Pagination{
			Page:          page,
			PageSize:      pageSize,
			TotalElements: totalElements,
		},
		Results: data,
	}
}

func BuildResponse(data any) ResponseNoPagination {
	return ResponseNoPagination{
		Results: data,
	}
}
