package customer

import (
	"context"
	"final-project/domain/entity"
)

func (c UsecaseCostumerInteractor) GetListCustomer(ctx context.Context) ([]*entity.Customer, error) {
	listCustomer, err := c.repoCustomer.GetListCustomer(ctx)
	if err != nil {
		return nil, err
	}

	// if include == "coupon" {
	// 	for _, customer := range listCustomer {
	// 		dataCoupon, errCou := c.repoCoupon.GetCouponById(ctx, customer.GetDataCoupon().GetCouponID())
	// 		if errCou != nil {
	// 			return nil, errCou
	// 		}
	// 		customer = customer.AddDataCoupon(dataCoupon)
	// 	}
	// }

	return listCustomer, nil
}
