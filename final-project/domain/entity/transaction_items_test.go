package entity_test

import (
	"final-project/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTransactionItems(t *testing.T) {
	transactionItem, err := entity.NewTransactionItems(entity.DTOTransactionItems{
		ItemID:        1,
		TransactionID: "TRAX18282828",
		CriteriaID:    2,
		RevenueItem:   500000,
	})
	now := time.Now()
	strDate := now.Format("2006-01-02 15:04:05")

	assert.Equal(t, 1, transactionItem.GetItemID())
	assert.Equal(t, "TRAX18282828", transactionItem.GetTransactionID())
	assert.Equal(t, 2, transactionItem.GetCriteriaID())
	assert.Equal(t, 500000, transactionItem.GetRevenueItem())
	assert.Equal(t, strDate, transactionItem.GetDateCreated())
	assert.Nil(t, err)
}
