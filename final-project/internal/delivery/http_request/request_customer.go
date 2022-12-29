package http_request

type RequestCustomer struct {
	CustomerID  string         `json:"customer_id"`
	Name        string         `json:"name"`
	Alamat      string         `json:"alamat"`
	PhoneNumber string         `json:"phone_number"`
	CreatedTime string         `json:"created_time"`
	Coupon      *RequestCoupon `json:"coupon"`
}
