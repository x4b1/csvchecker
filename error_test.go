package csvchecker

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToStringrowErrorFormatsCorrectly(t *testing.T) {
	const textError = "test error"
	testError := rowError{errors.New(textError)}

	assert.Equal(t, textError, testError.ToString())
}

func TestToStringcolErrorFormatsCorrectly(t *testing.T) {
	const textError, line, col = "test error", 22, 11
	testError := colError{
		line,
		col,
		errors.New(textError),
	}

	expectedError := fmt.Sprintf("Error on row %d and column %d: %s", line, col, textError)

	assert.Equal(t, expectedError, testError.ToString())
}
