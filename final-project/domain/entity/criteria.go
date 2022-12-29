package entity

import "errors"

type Criteria struct {
	criteriaID   int
	criteriaName string
}

type DTOCriteria struct {
	CriteriaID   int
	CriteriaName string
}

func NewCriteria(dto DTOCriteria) (*Criteria, error) {
	if dto.CriteriaID == 0 {
		return nil, errors.New("Criteria Cannot Be Empty")
	}
	if dto.CriteriaName == "" {
		return nil, errors.New("Criteria Name Cannot Be Empty")
	}

	criteria := &Criteria{
		criteriaID:   dto.CriteriaID,
		criteriaName: dto.CriteriaName,
	}
	return criteria, nil
}

func (c *Criteria) GetCriteriaID() int {
	return c.criteriaID
}

func (c *Criteria) GetCriteriaName() string {
	return c.criteriaName
}
