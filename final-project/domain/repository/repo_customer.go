package repository

import (
	"context"
	"final-project/domain/entity"
)

type InterfaceRepoCustomer interface {
	InsertDataCustomer(ctx context.Context, dataCustomer *entity.Customer) (string, error)
	GetListCustomer(ctx context.Context) ([]*entity.Customer, error)
	GetCustomerById(ctx context.Context, costumer_id string) (*entity.Customer, error)
}
