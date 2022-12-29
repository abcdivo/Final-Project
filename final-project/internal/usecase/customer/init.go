package customer

import "final-project/domain/repository"

type UsecaseCostumerInteractor struct {
	repoCustomer repository.InterfaceRepoCustomer
	repoCoupon   repository.InterfaceRepoCoupon
}

func NewCustomerUsecase(repoCostumer repository.InterfaceRepoCustomer, repoCoupon repository.InterfaceRepoCoupon) *UsecaseCostumerInteractor {
	return &UsecaseCostumerInteractor{
		repoCustomer: repoCostumer,
		repoCoupon:   repoCoupon,
	}
}
