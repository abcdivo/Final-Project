package coupon

import "final-project/domain/repository"

type UsecaseCouponInteractor struct {
	repoCoupon      repository.InterfaceRepoCoupon
	repoTransaction repository.InterfaceRepoTransaction
}

func NewCouponUsecase(repoCoupon repository.InterfaceRepoCoupon, repoTransaction repository.InterfaceRepoTransaction) *UsecaseCouponInteractor {
	return &UsecaseCouponInteractor{
		repoCoupon:      repoCoupon,
		repoTransaction: repoTransaction,
	}
}
