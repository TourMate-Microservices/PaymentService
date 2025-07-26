package request

type SearchPaginationRequest struct {
	Page       int    `json:"page" form:"page"`
	Keyword    string `json:"keyword" form:"keyword"`
	FilterProp string `json:"filterProp" form:"filterProp"` // Date, price, ...
	Order      string `json:"order" form:"order"`           //ASC or DESC
}
