package response

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	Data1    interface{}
	Data2    interface{}
	ErrMsg   error
	Context  *gin.Context
	PostType string
}

type MessageApiResponse struct {
	Message string `json:"message"`
}

type PaginationDataResponse struct {
	Data        interface{} `json:"data"`
	TotalCount  int         `json:"total_count"`
	Page        int         `json:"page"`
	PerPage     int         `json:"per_page"`
	TotalPages  int         `json:"total_pages"`
	HasNext     bool        `json:"has_next"`
	HasPrevious bool        `json:"has_previous"`
}
