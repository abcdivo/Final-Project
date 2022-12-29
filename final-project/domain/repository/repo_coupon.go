package repository

import (
	"context"
	"final-project/domain/entity"
)

type InterfaceRepoCoupon interface {
	InsertDataCoupon(ctx context.Context, dataCoupon *entity.Coupon) (string, error)
	GetCouponById(ctx context.Context, id_coupon string) (*entity.Coupon, error)
	CouponValidation(ctx context.Context, dataTransaction *entity.Transaction) (string, error)
	GetCouponByCustomerId(ctx context.Context, customer_id string) ([]*entity.Coupon, error)
	UpdateCouponStatus(ctx context.Context, coupon_id string, status int) error
	GetLastCreatedCouponByCustomerId(ctx context.Context, customer_id string) (*entity.Coupon, error)
}
