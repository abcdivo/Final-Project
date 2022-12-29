package coupon_response

import (
	"encoding/json"
	"final-project/domain/entity"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CustomReponseSingle struct {
	Status *Status             `json:"status"`
	Data   *ResponseCouponJson `json:"data"`
}
type CustomResponseCollection struct {
	Status *Status               `json:"status"`
	Data   []*ResponseCouponJson `json:"data"`
}

type ResponseCouponJson struct {
	CouponID    string `json:"coupon_id"`
	Types       string `json:"types"`
	ExpiredDate string `json:"expired_date"`
	CustomerID  string `json:"customer_id"`
	Status      int    `json:"status"`
}

type CouponID struct {
	CouponID string `json:"coupon_id"`
}

type ResponseSuccessInsert struct {
	Status *Status   `json:"status"`
	Data   *CouponID `json:"data"`
}

func MapResponseCoupon(dataCoupon *entity.Coupon, code int, message string, coupon_id string) ([]byte, error) {
	var resp *ResponseCouponJson
	if dataCoupon != nil {
		resp = &ResponseCouponJson{
			CouponID:    dataCoupon.GetCouponID(),
			Types:       dataCoupon.GetTypes(),
			ExpiredDate: dataCoupon.GetExpiredDate(),
			CustomerID:  dataCoupon.GetCustomerID(),
			Status:      dataCoupon.GetStatus(),
		}
	}

	httpResponse := &CustomReponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: resp,
	}

	if coupon_id != "" {
		httpResponse := &ResponseSuccessInsert{
			Status: &Status{
				Code:    code,
				Message: message,
			},
			Data: &CouponID{
				CouponID: coupon_id,
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
