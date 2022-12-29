package transaction_response

import (
	"encoding/json"
	"final-project/domain/entity"
)

type ResponseTransactionJson struct {
	TransactionID    string                      `json:"transaction_id"`
	CustomerID       string                      `json:"customer_id"`
	CustomerName     string                      `json:"customer_name"`
	Revenue          int                         `json:"revenue"`
	CouponID         string                      `json:"coupon_id"`
	PurchaseDate     string                      `json:"purchase_date"`
	TransactionItems []*ResponseTransactionItems `json:"transaction_items"`
}

type ResponseTransactionItems struct {
	TransactionID string `json:"transaction_id"`
	CriteriaID    int    `json:"criteria_id"`
	CriteriaName  string `json:"criteria_name"`
	RevenueItem   int    `json:"revenue_item"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TransactionID struct {
	TransactionID string `json:"transaction_id"`
}

type CustomReponseSingle struct {
	Status *Status                  `json:"status"`
	Data   *ResponseTransactionJson `json:"data"`
}

type ResponseSuccessInsert struct {
	Status *Status        `json:"status"`
	Data   *TransactionID `json:"data"`
}

type CustomReponseCollection struct {
	Status *Status                    `json:"status"`
	Data   []*ResponseTransactionJson `json:"data"`
}

func MapResponseTransaction(dataTransaction *entity.Transaction, code int, message string, transaction_id string) ([]byte, error) {
	var resp *ResponseTransactionJson
	if dataTransaction != nil {
		listItems := make([]*ResponseTransactionItems, 0)

		for _, item := range dataTransaction.GetTransactionItems() {
			transactionItem := &ResponseTransactionItems{
				TransactionID: item.GetTransactionID(),
				CriteriaID:    item.GetCriteriaID(),
				CriteriaName:  item.GetCriteriaName(),
				RevenueItem:   item.GetRevenueItem(),
			}

			listItems = append(listItems, transactionItem)
		}

		resp = &ResponseTransactionJson{
			TransactionID:    dataTransaction.GetTransactionID(),
			CustomerID:       dataTransaction.GetCustomerID(),
			CustomerName:     dataTransaction.GetCustomerName(),
			Revenue:          dataTransaction.GetRevenue(),
			CouponID:         dataTransaction.GetCouponID(),
			PurchaseDate:     dataTransaction.GetPurchaseDate(),
			TransactionItems: listItems,
		}
	}

	httpResponse := &CustomReponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: resp,
	}

	if transaction_id != "" {
		httpResponse := &ResponseSuccessInsert{
			Status: &Status{
				Code:    code,
				Message: message,
			},
			Data: &TransactionID{
				TransactionID: transaction_id,
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

func MapResponseListTransaction(listTransaction []*entity.Transaction, code int, message string) ([]byte, error) {
	listResp := make([]*ResponseTransactionJson, 0)
	for _, data := range listTransaction {
		listTransactionItem := make([]*ResponseTransactionItems, 0)
		for _, dataItem := range data.GetTransactionItems() {
			respItem := &ResponseTransactionItems{
				TransactionID: dataItem.GetTransactionID(),
				CriteriaID:    dataItem.GetCriteriaID(),
				CriteriaName:  dataItem.GetCriteriaName(),
				RevenueItem:   dataItem.GetRevenueItem(),
			}

			listTransactionItem = append(listTransactionItem, respItem)
		}

		resp := &ResponseTransactionJson{
			TransactionID:    data.GetTransactionID(),
			CustomerID:       data.GetCustomerID(),
			CustomerName:     data.GetCustomerName(),
			Revenue:          data.GetRevenue(),
			CouponID:         data.GetCouponID(),
			PurchaseDate:     data.GetPurchaseDate(),
			TransactionItems: listTransactionItem,
		}

		listResp = append(listResp, resp)
	}

	httpResponse := &CustomReponseCollection{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: listResp,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
