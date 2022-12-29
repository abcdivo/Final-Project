package customer_hendler

import (
	"context"
	"final-project/domain/repository"
	"final-project/domain/usecase"
	"final-project/internal/usecase/customer"
)

type CustomerHandler struct {
	customerUsecase usecase.CustomerService
	repoCustomer    repository.InterfaceRepoCustomer
	repoCoupon      repository.InterfaceRepoCoupon
	ctx             context.Context
}

func NewCustomerHandler(ctx context.Context, repoCustomer repository.InterfaceRepoCustomer, repoCoupon repository.InterfaceRepoCoupon) *CustomerHandler {
	customerUsecase := customer.NewCustomerUsecase(repoCustomer, repoCoupon)

	return &CustomerHandler{
		customerUsecase: customerUsecase,
		repoCustomer:    repoCustomer,
		repoCoupon:      repoCoupon,
		ctx:             ctx,
	}
}
