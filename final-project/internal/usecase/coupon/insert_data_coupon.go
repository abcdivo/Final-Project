package coupon

import (
	"context"
	"final-project/domain/entity"
)

func (c UsecaseCouponInteractor) InsertDataCoupon(ctx context.Context, dataCoupon *entity.Coupon) (string, error) {
	var dateLastCouponCreated string
	lastCreatedCoupon, errGetLastCoupon := c.repoCoupon.GetLastCreatedCouponByCustomerId(ctx, dataCoupon.GetCustomerID())
	if errGetLastCoupon != nil {
		return "", errGetLastCoupon
	}
	if lastCreatedCoupon != nil {
		dateLastCouponCreated = lastCreatedCoupon.GetDateCreated()
	}
	sumRevenue := c.repoTransaction.SumTransactionById(ctx, dataCoupon.GetCustomerID(), dateLastCouponCreated)
	_, errGenerateCouponId := dataCoupon.GenerateCouponId(sumRevenue)

	if errGenerateCouponId != nil {
		return "", errGenerateCouponId
	}

	if errGenerateCouponId != nil {
		return "", errGenerateCouponId
	}

	couponID, err := c.repoCoupon.InsertDataCoupon(ctx, dataCoupon)
	if err != nil {
		return "", nil
	}
	return couponID, nil
}
