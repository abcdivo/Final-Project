package customer_hendler

import (
	"encoding/json"
	"final-project/domain/entity"
	"final-project/internal/delivery/http_request"
	"final-project/internal/delivery/http_response/customer_response"
	"net/http"
	"time"
)

func (c *CustomerHandler) StoreDataCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestCustomer
		decoder = json.NewDecoder(r.Body)
	)

	errDecode := decoder.Decode(&req)

	if errDecode != nil {
		respErr, _ := customer_response.MapResponseCustomer(nil, http.StatusInternalServerError, "Error decode data", "")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respErr)
		return
	}

	var date = time.Now()

	customer, err := entity.NewCustomer(entity.DTOCustomer{
		Name:        req.Name,
		Alamat:      req.Alamat,
		PhoneNumber: req.PhoneNumber,
		CreatedTime: date.Format("2006-01-02"),
	})

	if err != nil {
		respErr, _ := customer_response.MapResponseCustomer(nil, http.StatusInternalServerError, err.Error(), "")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respErr)
		return
	}
	customer.SetUniqCustomerID()

	custID, errInsert := c.customerUsecase.InsertDataCustomer(c.ctx, customer)
	if errInsert != nil {
		respErr, _ := customer_response.MapResponseCustomer(nil, http.StatusInternalServerError, errInsert.Error(), "")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respErr)
		return
	}

	resp, errMap := customer_response.MapResponseCustomer(nil, http.StatusOK, "SUCCESS INSERT DATA", custID)
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
