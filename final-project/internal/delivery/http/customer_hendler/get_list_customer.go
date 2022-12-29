package customer_hendler

import (
	"context"
	"final-project/internal/delivery/http_response/customer_response"
	"net/http"
)

func (c *CustomerHandler) GetListCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.Background()
	)

	listCustomer, err := c.customerUsecase.GetListCustomer(ctx)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, errMap := customer_response.MapResponseListCustomer(listCustomer, 200, "Succes")
	if errMap != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Mapping data"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	w.Write(response)

	return
}
