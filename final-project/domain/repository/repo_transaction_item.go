package repository

import (
	"context"
	"final-project/domain/entity"
)

type InterfaceRepoTransactionItem interface {
	InsertDataTransactionItems(ctx context.Context, dataTransactionItems *entity.TransactionItems, transactionID string) error
	GetItemsByTransactionID(ctx context.Context, transaction_id string) ([]*entity.TransactionItems, error)
	DeleteItemsByTransactionId(ctx context.Context, transaction_id string) error
}
