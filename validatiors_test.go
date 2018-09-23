package csvchecker

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var validateStringErrorTests = []struct {
	validator  *StringValidation
	str        string
	expdErrStr string
}{
	{NewStringValidation(false, nil), "", "Value can't be empty"},
	{NewStringValidation(false, NewRangeValidation(0, 2)), "abcd", "Value abcd, maximun length 2, 4 given"},
	{NewStringValidation(false, NewRangeValidation(2, 6)), "a", "Value a, minimun length 2, 1 given"},
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
	validator *StringValidation
	str       string
}{
	{NewStringValidation(true, nil), ""},
	{NewStringValidation(false, NewRangeValidation(2, 2)), "ab"},
}

func TestValidateStringValidationReturnNil(t *testing.T) {
	for _, test := range validateStringCorrectTests {
		t.Run(test.str, func(t *testing.T) {
			assert.Nil(t, test.validator.Validate(test.str))
		})
	}
}

var validateNumberErrorTests = []struct {
	validator  *NumberValidation
	str        string
	expdErrStr string
}{
	{NewNumberValidation(false, nil), "asdb123", "Value asdb123, is not a number"},
	{NewNumberValidation(false, NewRangeValidation(0, 2)), "55", "Value 55, maximun value exceeded 2"},
	{NewNumberValidation(false, NewRangeValidation(-32, 2)), "-55", "Value -55, minimum value exceeded -32"},
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
	validator *NumberValidation
	str       string
}{
	{NewNumberValidation(false, nil), "32"},
	{NewNumberValidation(true, nil), ""},
	{NewNumberValidation(false, NewRangeValidation(2, 2)), "2"},
}

func TestValidateNumberValidationReturnNil(t *testing.T) {
	for _, test := range validateNumberCorrectTests {
		t.Run(test.str, func(t *testing.T) {
			assert.Nil(t, test.validator.Validate(test.str))
		})
	}
}

func TestValidateRegexpValidationReturnError(t *testing.T) {
	rgx := regexp.MustCompile("/[A-B]/i")
	v := NewRegexpValidation(rgx)

	assert.Error(t, v.Validate("Hello"))
}

func TestValidateRegexpValidationReturnNil(t *testing.T) {
	rgx := regexp.MustCompile("^correct.*.string$")
	v := NewRegexpValidation(rgx)

	assert.Nil(t, v.Validate("correct awesome string"))
}

var validateListValuesErrorTests = []struct {
	validator  *ListValuesValidator
	str        string
	expdErrStr string
}{
	{NewListValuesValidator(false, []string{}), "", "Value can't be empty"},
	{NewListValuesValidator(false, []string{"test", "test2"}), "test1", "Value is not in the list"},
}

func TestValidateListValuesValidationReturnError(t *testing.T) {
	for _, test := range validateListValuesErrorTests {
		t.Run(test.str, func(t *testing.T) {
			err := test.validator.Validate(test.str)
			if assert.Error(t, err) {
				assert.Equal(t, test.expdErrStr, err.Error())
			}
		})
	}
}

var validateListValuesCorrectTests = []struct {
	validator *ListValuesValidator
	str       string
}{
	{NewListValuesValidator(true, nil), ""},
	{NewListValuesValidator(false, []string{"test", "test2"}), "test"},
}

func TestValidateListValuesValidationReturnNil(t *testing.T) {
	for _, test := range validateListValuesCorrectTests {
		t.Run(test.str, func(t *testing.T) {
			assert.Nil(t, test.validator.Validate(test.str))
		})
	}
}
