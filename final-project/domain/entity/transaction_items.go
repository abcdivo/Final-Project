package entity

import (
	"errors"
	"time"
)

type TransactionItems struct {
	itemID        int
	transactionID string
	criteriaID    int
	criteriaName  string
	revenueItem   int
	dateCreated   time.Time
}

type DTOTransactionItems struct {
	ItemID        int
	TransactionID string
	CriteriaID    int
	CriteriaName  string
	RevenueItem   int
	DateCreated   string
}

func NewTransactionItems(dto DTOTransactionItems) (*TransactionItems, error) {
	// if dto.TransactionID == "" {
	// 	return nil, errors.New("transaction id cannot be empty")
	// }
	if dto.CriteriaID == 0 {
		return nil, errors.New("criteria id cannot be empty")
	}
	if dto.RevenueItem == 0 {
		return nil, errors.New("revenue cannot be empty")
	}

	transactionItems := &TransactionItems{
		itemID:        dto.ItemID,
		transactionID: dto.TransactionID,
		criteriaID:    dto.CriteriaID,
		criteriaName:  dto.CriteriaName,
		revenueItem:   dto.RevenueItem,
		dateCreated:   time.Now(),
	}

	return transactionItems, nil
}

func (ti *TransactionItems) GetItemID() int {
	return ti.itemID
}

func (ti *TransactionItems) GetTransactionID() string {
	return ti.transactionID
}

func (ti *TransactionItems) GetCriteriaID() int {
	return ti.criteriaID
}

func (ti *TransactionItems) GetRevenueItem() int {
	return ti.revenueItem
}

func (ti *TransactionItems) GetDateCreated() string {
	return ti.dateCreated.Format("2006-01-02 15:04:05")
}

func (ti *TransactionItems) GetCriteriaName() string {
	return ti.criteriaName
}
