package controllers

type ErrorResponse struct {
	Errors struct {
		Code    int         `json:"status_code"`
		Message string      `json:"message"`
		Errors  interface{} `json:"errors,omitempty"`
	} `json:"errors"`
}

type Errors struct {
	Code    int         `json:"status_code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

type PaginateResponse struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta,omitempty"`
}

type Meta struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Count       int   `json:"count"`
	CurrentPage int   `json:"current_page"`
	Links       Links `json:"links,omitempty"`
	PerPage     int   `json:"per_page"`
	Total       int   `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

type Links struct {
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
}
