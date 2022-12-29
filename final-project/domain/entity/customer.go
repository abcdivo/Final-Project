package entity

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type Customer struct {
	customerID  string
	name        string
	alamat      string
	phoneNumber string
	createdTime time.Time
	coupon      []*Coupon
}

type DTOCustomer struct {
	CustomerID  string
	Name        string
	Alamat      string
	PhoneNumber string
	CreatedTime string
	Coupon      []*Coupon
}

func NewCustomer(dto DTOCustomer) (*Customer, error) {
	if dto.Name == "" {
		return nil, errors.New("NAMA CANNOT BE EMPTY")
	}
	if dto.Alamat == "" {
		return nil, errors.New("ALAMAT CANNOT BE EMPTY")
	}
	if dto.PhoneNumber == "" {
		return nil, errors.New("PHONE NUMBER CANNOT BE EMPTY")
	}
	if dto.CreatedTime == "" {
		return nil, errors.New("CREATED TIME CANNOT BE EMPTY")
	}
	_, errConvertNumber := strconv.Atoi(dto.PhoneNumber)
	if errConvertNumber != nil {
		return nil, errors.New("PHONE NUMBER ONLY CONTAINS NUMERIC")
	}
	lenPhoneNumber := len(dto.PhoneNumber)
	if lenPhoneNumber != 12 {
		return nil, errors.New("PHONE NUMBER MUST 12 CHARACTER")
	}
	strCreatedTime, _ := time.Parse("2006-01-02", dto.CreatedTime)

	customer := &Customer{
		customerID:  dto.CustomerID,
		name:        dto.Name,
		alamat:      dto.Alamat,
		phoneNumber: dto.PhoneNumber,
		createdTime: strCreatedTime,
	}
	return customer, nil
}

// generate kode article
func (c *Customer) SetUniqCustomerID() *Customer {
	rand.Seed(time.Now().UnixNano())
	min := 10000000
	max := 30000000
	valueString := strconv.Itoa(rand.Intn(max-min+1) + min)
	c.customerID = "CUST" + valueString

	return c
}

func (c *Customer) AddDataCoupon(cou []*Coupon) *Customer {
	c.coupon = cou

	return c
}

func (c *Customer) GetCustomerID() string {
	return c.customerID
}

func (c *Customer) GetName() string {
	return c.name
}

func (c *Customer) GetAlamat() string {
	return c.alamat
}

func (c *Customer) GetPhoneNumber() string {
	return c.phoneNumber
}

func (c *Customer) GetCreatedTime() string {
	return c.createdTime.Format("2006-01-02")
}

func (c *Customer) GetDataCoupon() []*Coupon {
	return c.coupon
}
