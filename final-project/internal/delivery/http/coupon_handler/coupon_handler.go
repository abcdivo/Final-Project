package coupon_handler

import (
	"context"
	"final-project/domain/repository"
	"final-project/domain/usecase"
	"final-project/internal/usecase/coupon"
)

type CouponHandler struct {
	ctx           context.Context
	repoCoupon    repository.InterfaceRepoCoupon
	repoCustomer  repository.InterfaceRepoCustomer
	usecaseCoupon usecase.CouponService
}

func NewCouponHandler(ctx context.Context, repoCoupon repository.InterfaceRepoCoupon, repoCustomer repository.InterfaceRepoCustomer, repoTransaction repository.InterfaceRepoTransaction) *CouponHandler {
	usecaseCoupon := coupon.NewCouponUsecase(repoCoupon, repoTransaction)
	return &CouponHandler{
		ctx:           ctx,
		repoCoupon:    repoCoupon,
		repoCustomer:  repoCustomer,
		usecaseCoupon: usecaseCoupon,
	}
}

//
