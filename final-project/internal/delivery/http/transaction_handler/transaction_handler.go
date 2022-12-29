package transaction_handler

import (
	"context"
	"final-project/domain/repository"
	"final-project/domain/usecase"
	"final-project/internal/usecase/transaction"
)

type TransactionHandler struct {
	ctx                context.Context
	repoTransaction    repository.InterfaceRepoTransaction
	transactionUsecase usecase.TransactionService
}

func NewTransactionHandler(ctx context.Context, repoTransaction repository.InterfaceRepoTransaction, repoTransactionItem repository.InterfaceRepoTransactionItem, repoCoupon repository.InterfaceRepoCoupon) *TransactionHandler {
	transactionUsecas := transaction.NewTransactionUsecase(repoTransaction, repoTransactionItem, repoCoupon)
	return &TransactionHandler{
		repoTransaction:    repoTransaction,
		transactionUsecase: transactionUsecas,
		ctx:                ctx,
	}
}
