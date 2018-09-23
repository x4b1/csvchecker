package csvchecker

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Validator interface {
	Validate(string) error
}

type rangeValidation struct {
	min int
	max int
}

func NewRangeValidation(min, max int) *rangeValidation {
	return &rangeValidation{
		min: min,
		max: max,
	}
}

type stringValidation struct {
	allowEmpty bool
	inRange    *rangeValidation
}

func NewStringValidation(allowEmpty bool, inRange *rangeValidation) *stringValidation {
	return &stringValidation{
		allowEmpty: allowEmpty,
		inRange:    inRange,
	}
}

func (v *stringValidation) Validate(val string) error {
	strLn := len(val)
	if v.allowEmpty && len(val) < 1 {
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

type numberValidation struct {
	allowEmpty bool
	inRange    *rangeValidation
}

func NewNumberValidation(allowEmpty bool, inRange *rangeValidation) *numberValidation {
	return &numberValidation{
		allowEmpty: allowEmpty,
		inRange:    inRange,
	}
}

func (v *numberValidation) Validate(val string) error {

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

type regexpValidation struct {
	regexp *regexp.Regexp
}

func NewRegexpValidation(regexp *regexp.Regexp) *regexpValidation {
	return &regexpValidation{regexp: regexp}
}

func (v *regexpValidation) Validate(val string) error {
	if !v.regexp.MatchString(val) {
		return fmt.Errorf("Value %s, doesn't match with %s regex", val, v.regexp.String())
	}

	return nil
}
