package csvchecker

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// Validator is an interface that should implements all validators
type Validator interface {
	Validate(string) error
}

// RangeValidation represents de minimun and maximun values in a validation
type RangeValidation struct {
	min int
	max int
}

// NewRangeValidation creates a new instance of RangeValidation
func NewRangeValidation(min, max int) *RangeValidation {
	return &RangeValidation{
		min: min,
		max: max,
	}
}

// StringValidation is the struct to check string fields
type StringValidation struct {
	allowEmpty bool
	inRange    *RangeValidation
}

// NewStringValidation creates a new instance of StringValidation
func NewStringValidation(allowEmpty bool, inRange *RangeValidation) *StringValidation {
	return &StringValidation{
		allowEmpty: allowEmpty,
		inRange:    inRange,
	}
}

// Validate applies the rules configured in validator
func (v *StringValidation) Validate(val string) error {
	strLn := len(val)
	if v.allowEmpty && strLn < 1 {
		return nil
	} else if !v.allowEmpty && strLn < 1 {
		return errors.New("Value can't be empty")
	}

	if v.inRange != nil && v.inRange.max < strLn {
		return fmt.Errorf("Value %s, maximun length %d, %d given", val, v.inRange.max, strLn)
	}

	if v.inRange != nil && v.inRange.min > strLn {
		return fmt.Errorf("Value %s, minimun length %d, %d given", val, v.inRange.min, strLn)
	}

	return nil
}

// NumberValidation is the struct to check number fields
type NumberValidation struct {
	allowEmpty bool
	inRange    *RangeValidation
}

// NewNumberValidation creates a new instance of NumberValidation
func NewNumberValidation(allowEmpty bool, inRange *RangeValidation) *NumberValidation {
	return &NumberValidation{
		allowEmpty: allowEmpty,
		inRange:    inRange,
	}
}

// Validate applies the rules configured in validator
func (v *NumberValidation) Validate(val string) error {

	if v.allowEmpty && len(val) < 1 {
		return nil
	}

	nVal, err := strconv.Atoi(val)

	if err != nil {
		return fmt.Errorf("Value %s, is not a number", val)
	}

	if v.inRange != nil && v.inRange.max < nVal {
		return fmt.Errorf("Value %s, maximun value exceeded %d", val, v.inRange.max)
	}

	if v.inRange != nil && v.inRange.min > nVal {
		return fmt.Errorf("Value %s, minimum value exceeded %d", val, v.inRange.min)
	}

	return nil
}

// RegexpValidation is the struct to check fields given a regexp
type RegexpValidation struct {
	regexp *regexp.Regexp
}

// NewRegexpValidation creates a new instance of RegexpValidation
func NewRegexpValidation(regexp *regexp.Regexp) *RegexpValidation {
	return &RegexpValidation{regexp: regexp}
}

// Validate applies the rules configured in validator
func (v *RegexpValidation) Validate(val string) error {
	if !v.regexp.MatchString(val) {
		return fmt.Errorf("Value %s, doesn't match with %s regex", val, v.regexp.String())
	}

	return nil
}

// ListValuesValidator is the struct to check string fields that must match with given values
type ListValuesValidator struct {
	allowEmpty bool
	values     []string
}

// NewListValuesValidator creates a new instance of ListValuesValidator
func NewListValuesValidator(allowEmpty bool, values []string) *ListValuesValidator {
	return &ListValuesValidator{
		allowEmpty: allowEmpty,
		values:     values,
	}
}

// Validate applies the rules configured in validator
func (v *ListValuesValidator) Validate(val string) error {
	strLn := len(val)
	if v.allowEmpty && strLn < 1 {
		return nil
	} else if !v.allowEmpty && strLn < 1 {
		return errors.New("Value can't be empty")
	}

	for _, value := range v.values {
		if val == value {
			return nil
		}
	}

	return errors.New("Value is not in the list")
}
