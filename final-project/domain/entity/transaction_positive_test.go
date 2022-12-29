package entity_test

import (
	"final-project/domain/entity"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestScenarioRevenueWithTypes struct {
	Types            string
	DiscountExpected int
	Want             int
}

func TestNewTransactionWithoutTypes(t *testing.T) {
	var types string
	transaction, err := entity.NewTransaction(entity.DTOTransaction{
		CustomerID:   "CUST-001",
		CouponID:     "",
		PurchaseDate: "2022-11-28",
	})
	transaction.SetUniqTransactionID()

	transactionItems := make([]*entity.TransactionItems, 0)
	items1, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		TransactionID: transaction.GetTransactionID(),
		CriteriaID:    2,
		RevenueItem:   1000000,
	})
	items2, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		TransactionID: transaction.GetTransactionID(),
		CriteriaID:    1,
		RevenueItem:   500000,
	})
	items3, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		TransactionID: transaction.GetTransactionID(),
		CriteriaID:    1,
		RevenueItem:   250000,
	})
	items4, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		TransactionID: transaction.GetTransactionID(),
		CriteriaID:    5,
		RevenueItem:   150000,
	})
	transactionItems = append(transactionItems, items1, items2, items3, items4)
	transaction.SetTransactionItems(transactionItems)
	types = ""
	total_revenue, _ := transaction.SumTotalRevenue(types)

	fmt.Println(total_revenue)
	fmt.Println(transaction.GetRevenue())
	assert.Equal(t, "CUST-001", transaction.GetCustomerID())
	assert.Equal(t, 1900000, transaction.GetRevenue())
	assert.Equal(t, "2022-11-28", transaction.GetPurchaseDate())
	assert.Nil(t, err)
}

func TestNewTransactionWithTypes(t *testing.T) {
	transaction, err := entity.NewTransaction(entity.DTOTransaction{
		CustomerID:   "CUST-001",
		CouponID:     "",
		PurchaseDate: "2022-11-28",
	})
	transaction.SetUniqTransactionID()

	transactionItems := make([]*entity.TransactionItems, 0)
	items1, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		TransactionID: transaction.GetTransactionID(),
		CriteriaID:    2,
		RevenueItem:   1000000,
	})
	items2, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		TransactionID: transaction.GetTransactionID(),
		CriteriaID:    1,
		RevenueItem:   500000,
	})
	items3, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		TransactionID: transaction.GetTransactionID(),
		CriteriaID:    1,
		RevenueItem:   250000,
	})
	items4, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		TransactionID: transaction.GetTransactionID(),
		CriteriaID:    5,
		RevenueItem:   150000,
	})
	items5, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		TransactionID: transaction.GetTransactionID(),
		CriteriaID:    3,
		RevenueItem:   500000,
	})
	items6, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		TransactionID: transaction.GetTransactionID(),
		CriteriaID:    4,
		RevenueItem:   300000,
	})
	transactionItems = append(transactionItems, items1, items2, items3, items4, items5, items6)
	transaction.SetTransactionItems(transactionItems)

	listScenario := []TestScenarioRevenueWithTypes{
		{
			Types:            "ULTI",
			DiscountExpected: 195000,
			Want:             2505000,
		},
		{
			Types:            "PREMI",
			DiscountExpected: 112500,
			Want:             2587500,
		},
		{
			Types:            "BASIC",
			DiscountExpected: 15000,
			Want:             2685000,
		},
	}
	for _, data := range listScenario {
		total_revenue, _ := transaction.SumTotalRevenue(data.Types)

		fmt.Println(total_revenue)
		fmt.Println(transaction.GetRevenue())
		assert.Equal(t, "CUST-001", transaction.GetCustomerID())
		// assert.Equal(t, 1787500, transaction.GetRevenue())
		assert.Equal(t, data.Want, transaction.GetRevenue())
		assert.Equal(t, "2022-11-28", transaction.GetPurchaseDate())
		assert.Nil(t, err)
	}
}
