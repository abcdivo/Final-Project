package customer_response

import (
	"encoding/json"
	"final-project/domain/entity"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CustomReponseSingle struct {
	Status *Status           `json:"status"`
	Data   *ResponseCustomer `json:"data"`
}
type CustomResponseCollection struct {
	Status *Status                 `json:"status"`
	Data   []*ResponseCustomerList `json:"data"`
}

type ResponseCustomer struct {
	CustomerID  string            `json:"customer_id"`
	Name        string            `json:"name"`
	Alamat      string            `json:"alamat"`
	PhoneNumber string            `json:"phone_number"`
	CreatedTime string            `json:"created_time"`
	Coupon      []*ResponseCoupon `json:"coupon"`
}

type ResponseCustomerList struct {
	CustomerID  string `json:"customer_id"`
	Name        string `json:"name"`
	Alamat      string `json:"alamat"`
	PhoneNumber string `json:"phone_number"`
	CreatedTime string `json:"created_time"`
}

type ResponseCoupon struct {
	CouponID    string `json:"coupon_id"`
	Types       string `json:"types"`
	ExpiredDate string `json:"expired_date"`
	CustomerID  string `json:"customer_id"`
	Status      int    `json:"status"`
}

type CustomerID struct {
	CustomerID string `json:"customer_id"`
}

type ResponseSuccessInsert struct {
	Status *Status     `json:"status"`
	Data   *CustomerID `json:"data"`
}

func MapResponseCustomer(dataCustomer *entity.Customer, code int, message string, customer_id string) ([]byte, error) {
	var resp *ResponseCustomer
	if dataCustomer != nil {
		listCoupon := make([]*ResponseCoupon, 0)
		for _, dataCoupon := range dataCustomer.GetDataCoupon() {
			responseCoupon := &ResponseCoupon{
				CouponID:    dataCoupon.GetCouponID(),
				Types:       dataCoupon.GetTypes(),
				ExpiredDate: dataCoupon.GetExpiredDate(),
				CustomerID:  dataCoupon.GetCustomerID(),
				Status:      dataCoupon.GetStatus(),
			}
			listCoupon = append(listCoupon, responseCoupon)
		}

		resp = &ResponseCustomer{
			CustomerID:  dataCustomer.GetCustomerID(),
			Name:        dataCustomer.GetName(),
			Alamat:      dataCustomer.GetAlamat(),
			PhoneNumber: dataCustomer.GetPhoneNumber(),
			CreatedTime: dataCustomer.GetCreatedTime(),
			Coupon:      listCoupon,
		}
	}

	httpResponse := &CustomReponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: resp,
	}

	if customer_id != "" {
		httpResponse := &ResponseSuccessInsert{
			Status: &Status{
				Code:    code,
				Message: message,
			},
			Data: &CustomerID{
				CustomerID: customer_id,
			},
		}
		respJson, err := json.Marshal(httpResponse)
		if err != nil {
			return nil, err
		}

		return respJson, nil
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

func MapResponseListCustomer(listCustomer []*entity.Customer, code int, message string) ([]byte, error) {
	lisResp := make([]*ResponseCustomerList, 0)
	for _, data := range listCustomer {
		rsp := &ResponseCustomerList{
			CustomerID:  data.GetCustomerID(),
			Name:        data.GetName(),
			Alamat:      data.GetAlamat(),
			PhoneNumber: data.GetPhoneNumber(),
			CreatedTime: data.GetCreatedTime(),
		}
		lisResp = append(lisResp, rsp)
	}
	httpResponse := &CustomResponseCollection{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: lisResp,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
