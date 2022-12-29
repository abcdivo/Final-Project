package transaction

import "final-project/domain/repository"

type UsecaseTransactionInteractor struct {
	repoTransaction      repository.InterfaceRepoTransaction
	repoTransactionItems repository.InterfaceRepoTransactionItem
	repoCoupon           repository.InterfaceRepoCoupon
}

func NewTransactionUsecase(repoTransaction repository.InterfaceRepoTransaction, repoTransactionItems repository.InterfaceRepoTransactionItem, repoCoupon repository.InterfaceRepoCoupon) *UsecaseTransactionInteractor {
	return &UsecaseTransactionInteractor{
		repoTransaction:      repoTransaction,
		repoTransactionItems: repoTransactionItems,
		repoCoupon:           repoCoupon,
	}
}
