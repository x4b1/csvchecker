package csvchecker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumnCreatedWithPositionAndWithOutValidator(t *testing.T) {
	pos := 2
	dummyValidator := new(StringValidation)
	column := NewColumn(pos, dummyValidator)
	assert.Equal(t, pos, column.position)
	assert.IsType(t, dummyValidator, column.validator)
}
