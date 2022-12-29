package mapper

import "final-project/domain/entity"

func DataCustomerDbToEntity(dataDTO entity.DTOCustomer) (*entity.Customer, error) {
	customer, err := entity.NewCustomer(dataDTO)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
