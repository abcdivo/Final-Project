package entity_test

import (
	"final-project/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCriteriaValidation struct {
	CriteriaID   int
	CriteriaName string
}

func TestNewCriteriaValidation(t *testing.T) {
	dataCriteria, err := entity.NewCriteria(entity.DTOCriteria{
		CriteriaID:   2,
		CriteriaName: "Joni",
	})

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 2, dataCriteria.GetCriteriaID())
	assert.Equal(t, "Joni", dataCriteria.GetCriteriaName())
}
