package transaction_handler

import (
	"context"
	"final-project/internal/delivery/http_response/transaction_response"
	"net/http"
	"strconv"
)

func (tr *TransactionHandler) GetListTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.Background()
	)

	params := r.URL.Query().Get("limit")
	limit, errParse := strconv.Atoi(params)
	if errParse != nil {
		limit = 10
	}

	listTransaction, err := tr.transactionUsecase.GetListTransaction(ctx, limit)
	if err != nil {
		errResp, _ := transaction_response.MapResponseListTransaction(nil, http.StatusInternalServerError, err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errResp)
		return
	}

	response, errMap := transaction_response.MapResponseListTransaction(listTransaction, http.StatusOK, "Success")
	if errMap != nil {
		errResp, _ := transaction_response.MapResponseListTransaction(nil, http.StatusInternalServerError, errMap.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}
