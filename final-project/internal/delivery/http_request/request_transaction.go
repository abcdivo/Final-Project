package http_request

type RequestTransaction struct {
	CustomerID string `json:"customer_id"`
	ProductID  string `json:"product_id"`
	Quantity   int    `json:"quantity"`
	// Revenue          int                       `json:"revenue"`
	CouponID         string                    `json:"coupon_id"`
	PurchaseDate     string                    `json:"purchase_date"`
	TransactionItems []*RequestTransactionItem `json:"transaction_items"`
}
