package entity_test

import (
	"final-project/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCustomerValidation struct {
	UserID      int
	Name        string
	Alamat      string
	PhoneNumber string
	CreatedTime string
}

func TestNewDataCustomerValidation(t *testing.T) {
	dataCustomer, err := entity.NewCustomer(entity.DTOCustomer{
		Name:        "Divo",
		Alamat:      "Padang",
		PhoneNumber: "0821756563",
		CreatedTime: "2022-11-27",
	})

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Divo", dataCustomer.GetName())
	assert.Equal(t, "Padang", dataCustomer.GetAlamat())
	assert.Equal(t, "0821756563", dataCustomer.GetPhoneNumber())
	assert.Equal(t, "2022-11-27", dataCustomer.GetCreatedTime())

}
