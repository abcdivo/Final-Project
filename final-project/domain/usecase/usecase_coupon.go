package usecase

import (
	"context"
	"final-project/domain/entity"
)

type CouponService interface {
	InsertDataCoupon(ctx context.Context, dataCoupon *entity.Coupon) (string, error)
}
