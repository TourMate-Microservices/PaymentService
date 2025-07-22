package request

type SearchPaginationRequest struct {
	Page       int    `json:"page" form:"page"`
	Keyword    string `json:"keyword" form:"keyword"`
	FilterProp string `json:"filter_prop" form:"filter_prop"` // Date, price, ...
	Order      string `json:"order" form:"order"`             //ASC or DESC
}
