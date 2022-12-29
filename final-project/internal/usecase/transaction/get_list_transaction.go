package transaction

import (
	"context"
	"final-project/domain/entity"
)

func (tr UsecaseTransactionInteractor) GetListTransaction(ctx context.Context, limit int) ([]*entity.Transaction, error) {
	listTransaction, err := tr.repoTransaction.GetListTransaction(ctx, limit)
	for _, data := range listTransaction {
		items, errTransactionItem := tr.repoTransactionItems.GetItemsByTransactionID(ctx, data.GetTransactionID())
		data.SetTransactionItems(items)
		if errTransactionItem != nil {
			return nil, errTransactionItem
		}
	}
	if err != nil {
		return nil, err
	}

	return listTransaction, nil
}
