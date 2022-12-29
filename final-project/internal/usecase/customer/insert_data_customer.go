package customer

import (
	"context"
	"final-project/domain/entity"
)

func (c UsecaseCostumerInteractor) InsertDataCustomer(ctx context.Context, dataCustomer *entity.Customer) (string, error) {
	custID, err := c.repoCustomer.InsertDataCustomer(ctx, dataCustomer)
	if err != nil {
		return "", err
	}

	return custID, err
}
