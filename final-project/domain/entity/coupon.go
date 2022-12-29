package entity

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Coupon struct {
	couponID    string
	types       string
	expiredDate time.Time
	customerID  string
	status      int
	createdDate time.Time
}

type DTOCoupon struct {
	CouponID    string
	Types       string
	ExpiredDate string
	CustomerID  string
	Status      int
	CreatedDate string
}

func NewCoupon(dto DTOCoupon) (*Coupon, error) {
	if dto.CustomerID == "" {
		return nil, errors.New("customer id cannot be empty")
	}

	today := time.Now()
	expiredDate := today.AddDate(1, 0, 0)

	coupon := &Coupon{
		couponID:    dto.CouponID,
		types:       dto.Types,
		expiredDate: expiredDate,
		customerID:  dto.CustomerID,
		status:      dto.Status,
		createdDate: time.Now(),
	}

	return coupon, nil
}

func (c *Coupon) GetCouponID() string {
	return c.couponID
}

func (c *Coupon) GetTypes() string {
	return c.types
}
func (c *Coupon) GetCustomerID() string {
	return c.customerID
}

func (c *Coupon) GetStatus() int {
	return c.status
}

func (c *Coupon) GetExpiredDate() string {
	return c.expiredDate.Format("2006-01-02")
}

func (c *Coupon) GetDateCreated() string {
	return c.createdDate.Format("2006-01-02 15:04:05")
}

func (c *Coupon) ValidateCouponTypes(sumRevenue int64) string {
	var typesString string
	if sumRevenue > int64(25000000) {
		typesString = "ULTI"
	} else if sumRevenue > int64(13000000) {
		typesString = "PREMI"
	} else if sumRevenue > int64(6000000) {
		typesString = "BASIC"
	}

	c.types = typesString
	return typesString
}

func (c *Coupon) GenerateCouponId(sumRevenue int64) (string, error) {
	rand.Seed(time.Now().UnixNano())
	min := 1000000000000
	max := 9999999999999
	valueString := strconv.Itoa(rand.Intn(max-min+1) + min)
	typesString := c.ValidateCouponTypes(sumRevenue)
	if typesString == "" {
		errString := fmt.Sprintf("this customer did not meet the criteria. Total Revenue Customer = %d", sumRevenue)
		return "", errors.New(errString)
	}
	c.couponID = typesString + "-RND" + valueString

	return c.couponID, nil
}
