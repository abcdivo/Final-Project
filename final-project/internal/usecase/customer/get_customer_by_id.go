package customer

import (
	"context"
	"final-project/domain/entity"
)

func (c UsecaseCostumerInteractor) GetCustomerById(ctx context.Context, customer_id string) (*entity.Customer, error) {

	listCoupon, errcou := c.repoCoupon.GetCouponByCustomerId(ctx, customer_id)
	if errcou != nil {
		return nil, errcou
	}
	listCustomer, err := c.repoCustomer.GetCustomerById(ctx, customer_id)
	if err != nil {
		return nil, err
	}
	listCustomer.AddDataCoupon(listCoupon)

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
