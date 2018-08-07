package csvchecker

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStringValidationWithDefaultValues(t *testing.T) {
	v := StringValidation{}
	assert.False(t, v.AllowEmpty)
	assert.Nil(t, v.InRange)
}

var validateStringErrorTests = []struct {
	validator  StringValidation
	str        string
	expdErrStr string
}{
	{StringValidation{}, "", "Value can't be empty"},
	{StringValidation{false, &RangeValidation{0, 2}}, "abcd", "Value abcd, maximun length 2, 4 given"},
	{StringValidation{false, &RangeValidation{2, 6}}, "a", "Value a, minimun length 2, 1 given"},
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
	validator StringValidation
	str       string
}{
	{StringValidation{AllowEmpty: true}, ""},
	{StringValidation{false, &RangeValidation{2, 2}}, "ab"},
}

func TestValidateStringValidationReturnNil(t *testing.T) {
	for _, test := range validateStringCorrectTests {
		t.Run(test.str, func(t *testing.T) {
			assert.Nil(t, test.validator.Validate(test.str))
		})
	}
}

func TestCreateNumberValidationWithDefaultValues(t *testing.T) {
	v := NumberValidation{}
	assert.Nil(t, v.InRange)
}

var validateNumberErrorTests = []struct {
	validator  NumberValidation
	str        string
	expdErrStr string
}{
	{NumberValidation{}, "asdb123", "Value asdb123, is not a number"},
	{NumberValidation{AllowEmpty: false, InRange: &RangeValidation{0, 2}}, "55", "Value 55, maximun value exceeded 2"},
	{NumberValidation{AllowEmpty: false, InRange: &RangeValidation{-32, 2}}, "-55", "Value -55, minimum value exceeded -32"},
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
	validator NumberValidation
	str       string
}{
	{NumberValidation{}, "32"},
	{NumberValidation{AllowEmpty: true}, ""},
	{NumberValidation{AllowEmpty: false, InRange: &RangeValidation{2, 2}}, "2"},
}

func TestValidateNumberValidationReturnNil(t *testing.T) {
	for _, test := range validateNumberCorrectTests {
		t.Run(test.str, func(t *testing.T) {
			assert.Nil(t, test.validator.Validate(test.str))
		})
	}
}

func TestCreateRegexValidationWithDefaultValues(t *testing.T) {
	v := RegexValidation{}
	assert.IsType(t, regexp.Regexp{}, v.Regex)
}

func TestValidateRegexValidationReturnError(t *testing.T) {
	rgx := regexp.MustCompile("/[A-B]/i")
	v := RegexValidation{*rgx}

	assert.Error(t, v.Validate("Hello"))
}

func TestValidateRegexValidationReturnNil(t *testing.T) {
	rgx := regexp.MustCompile("^correct.*.string$")
	v := RegexValidation{*rgx}

	assert.Nil(t, v.Validate("correct awesome string"))
}
