package customer_hendler

import (
	"context"
	"final-project/internal/delivery/http_response/customer_response"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *CustomerHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.Background()
	)
	vars := mux.Vars(r)
	id := vars["id"]

	customer, err := c.customerUsecase.GetCustomerById(ctx, id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, errMap := customer_response.MapResponseCustomer(customer, 200, "Succes", "")
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
