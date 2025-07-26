package response

type PaymentCallbackComponent struct {
	//PaymentType   string  `json:"paymentType" form:"paymentType"` // bo^' bo? theo y' m
	CustomerId    int     `json:"customerId" form:"customerId"`
	AccountId     int     `json:"accountId" form:"accountId"`
	PaymentMethod string  `json:"paymentMethod" form:"paymentMethod"`
	Price         float64 `json:"price" form:"price"`
	OrderCode     int     `json:"orderCode" form:"orderCode"`
}
