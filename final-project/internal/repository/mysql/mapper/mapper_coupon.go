package mapper

import "final-project/domain/entity"

func DataCouponDbToEntity(dataDTO entity.DTOCoupon) (*entity.Coupon, error) {
	coupon, err := entity.NewCoupon(dataDTO)
	if err != nil {
		return nil, err
	}

	return coupon, nil
}
