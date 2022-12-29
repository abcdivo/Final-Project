package transaction_handler

import (
	"encoding/json"
	"final-project/domain/entity"
	"final-project/internal/delivery/http_request"
	"final-project/internal/delivery/http_response/transaction_response"
	"net/http"
)

func (tr *TransactionHandler) StoreDataTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestTransaction
		decoder = json.NewDecoder(r.Body)
	)
	errDecode := decoder.Decode(&req)

	if errDecode != nil {
		respErr, _ := transaction_response.MapResponseTransaction(nil, http.StatusInternalServerError, "Error decode data", "")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respErr)
		return
	}

	listItems := make([]*entity.TransactionItems, 0)

	for _, item := range req.TransactionItems {
		transactionItem, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
			CriteriaID:  item.CriteriaID,
			RevenueItem: item.RevenueItem,
		})

		listItems = append(listItems, transactionItem)
	}

	transaction, err := entity.NewTransaction(entity.DTOTransaction{
		CustomerID: req.CustomerID,
		// Revenue:          req.Revenue,
		CouponID:         req.CouponID,
		PurchaseDate:     req.PurchaseDate,
		TransactionItems: listItems,
	})
	if err != nil {
		respErr, _ := transaction_response.MapResponseTransaction(nil, http.StatusInternalServerError, err.Error(), "")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respErr)
		return
	}
	transaction.SetUniqTransactionID()
	traxID, errInsert := tr.transactionUsecase.InsertDataTransaction(tr.ctx, transaction)
	if errInsert != nil {
		respErr, _ := transaction_response.MapResponseTransaction(nil, http.StatusInternalServerError, errInsert.Error(), "")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respErr)
		return
	}

	resp, errMap := transaction_response.MapResponseTransaction(nil, http.StatusOK, "SUCCESS INSERT DATA", traxID)
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
