package csvchecker

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var validateStringErrorTests = []struct {
	validator  *stringValidation
	str        string
	expdErrStr string
}{
	{NewStringValidation(false, nil), "", "Value can't be empty"},
	{NewStringValidation(false, newRangeValidation(0, 2)), "abcd", "Value abcd, maximun length 2, 4 given"},
	{NewStringValidation(false, newRangeValidation(2, 6)), "a", "Value a, minimun length 2, 1 given"},
}

func TestValidateStringValidationReturnError(t *testing.T) {
	for _, test := range validateStringErrorTests {
		t.Run(test.str, func(t *testing.T) {
			err := test.validator.Validate(test.str)
			if assert.Error(t, err) {
				assert.Equal(t, test.expdErrStr, err.Error())
			}
		})
	}
}

var validateStringCorrectTests = []struct {
	validator *stringValidation
	str       string
}{
	{NewStringValidation(true, nil), ""},
	{NewStringValidation(false, newRangeValidation(2, 2)), "ab"},
}

func TestValidateStringValidationReturnNil(t *testing.T) {
	for _, test := range validateStringCorrectTests {
		t.Run(test.str, func(t *testing.T) {
			assert.Nil(t, test.validator.Validate(test.str))
		})
	}
}

var validateNumberErrorTests = []struct {
	validator  *numberValidation
	str        string
	expdErrStr string
}{
	{NewNumberValidation(false, nil), "asdb123", "Value asdb123, is not a number"},
	{NewNumberValidation(false, newRangeValidation(0, 2)), "55", "Value 55, maximun value exceeded 2"},
	{NewNumberValidation(false, newRangeValidation(-32, 2)), "-55", "Value -55, minimum value exceeded -32"},
}

func TestValidateNumberValidationReturnError(t *testing.T) {
	for _, test := range validateNumberErrorTests {
		t.Run(test.str, func(t *testing.T) {
			err := test.validator.Validate(test.str)
			if assert.Error(t, err) {
				assert.Equal(t, test.expdErrStr, err.Error())
			}
		})
	}
}

var validateNumberCorrectTests = []struct {
	validator *numberValidation
	str       string
}{
	{NewNumberValidation(false, nil), "32"},
	{NewNumberValidation(true, nil), ""},
	{NewNumberValidation(false, newRangeValidation(2, 2)), "2"},
}

func TestValidateNumberValidationReturnNil(t *testing.T) {
	for _, test := range validateNumberCorrectTests {
		t.Run(test.str, func(t *testing.T) {
			assert.Nil(t, test.validator.Validate(test.str))
		})
	}
}

func TestValidateregexpValidationReturnError(t *testing.T) {
	rgx := regexp.MustCompile("/[A-B]/i")
	v := NewRegexpValidation(rgx)

	assert.Error(t, v.Validate("Hello"))
}

func TestValidateregexpValidationReturnNil(t *testing.T) {
	rgx := regexp.MustCompile("^correct.*.string$")
	v := NewRegexpValidation(rgx)

	assert.Nil(t, v.Validate("correct awesome string"))
}
