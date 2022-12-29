package entity_test

import (
	"final-project/domain/entity"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCoupon(t *testing.T) {
	coupon, err := entity.NewCoupon(entity.DTOCoupon{
		CustomerID: "CUST15399180",
	})

	fmt.Println(coupon.GetExpiredDate())
	assert.Equal(t, "2023-12-23", coupon.GetExpiredDate())
	assert.Equal(t, "CUST15399180", coupon.GetCustomerID())
	assert.Equal(t, 0, coupon.GetStatus())
	assert.Nil(t, err)
}

func TestGenerateCouponId(t *testing.T) {
	// sumRevenue := 4000000
	coupon, _ := entity.NewCoupon(entity.DTOCoupon{
		CustomerID: "CUST15399180",
	})

	couponID, errGenerateCouponId := coupon.GenerateCouponId(10000000)
	fmt.Println(couponID)
	assert.Nil(t, errGenerateCouponId)
}
