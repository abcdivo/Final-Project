package coupon_handler

import (
	"encoding/json"
	"final-project/domain/entity"
	"final-project/internal/delivery/http_request"
	"final-project/internal/delivery/http_response/coupon_response"
	"net/http"
)

func (c *CouponHandler) StoreDataCoupon(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestCoupon
		decoder = json.NewDecoder(r.Body)
	)

	errDecode := decoder.Decode(&req)

	if errDecode != nil {
		respErr, _ := coupon_response.MapResponseCoupon(nil, http.StatusInternalServerError, "Error decode data", "")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respErr)
		return
	}

	coupon, err := entity.NewCoupon(entity.DTOCoupon{
		CustomerID: req.CustomerID,
	})

	if err != nil {
		respErr, _ := coupon_response.MapResponseCoupon(nil, http.StatusInternalServerError, err.Error(), "")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respErr)
		return
	}

	coupon_id, errInsert := c.usecaseCoupon.InsertDataCoupon(c.ctx, coupon)
	if errInsert != nil {
		respErr, _ := coupon_response.MapResponseCoupon(nil, http.StatusInternalServerError, errInsert.Error(), "")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respErr)
		return
	}

	resp, errMap := coupon_response.MapResponseCoupon(nil, http.StatusOK, "SUCCESS INSERT DATA", coupon_id)
	if errMap != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMap.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return
}
