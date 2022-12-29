package usecase

import (
	"context"
	"final-project/domain/entity"
)

type TransactionService interface {
	InsertDataTransaction(ctx context.Context, dataTransaction *entity.Transaction) (string, error)
	GetListTransaction(ctx context.Context, limit int) ([]*entity.Transaction, error)
}
