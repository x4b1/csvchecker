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

type RangeValidation struct {
	Min int
	Max int
}

type StringValidation struct {
	AllowEmpty bool
	InRange    *RangeValidation
}

func (v *StringValidation) Validate(val string) error {
	strLn := len(val)
	if v.AllowEmpty && len(val) < 1 {
		return nil
	} else if !v.AllowEmpty && strLn < 1 {
		return errors.New("Value can't be empty")
	}

	if v.InRange != nil && v.InRange.Max < strLn {
		return fmt.Errorf("Value %s, maximun length %d, %d given", val, v.InRange.Max, strLn)
	}

	if v.InRange != nil && v.InRange.Min > strLn {
		return fmt.Errorf("Value %s, minimun length %d, %d given", val, v.InRange.Min, strLn)
	}

	return nil
}

type NumberValidation struct {
	AllowEmpty bool
	InRange    *RangeValidation
}

func (v *NumberValidation) Validate(val string) error {

	if v.AllowEmpty && len(val) < 1 {
		return nil
	}

	nVal, err := strconv.Atoi(val)

	if err != nil {
		return fmt.Errorf("Value %s, is not a number", val)
	}

	if v.InRange != nil && v.InRange.Max < nVal {
		return fmt.Errorf("Value %s, maximun value exceeded %d", val, v.InRange.Max)
	}

	if v.InRange != nil && v.InRange.Min > nVal {
		return fmt.Errorf("Value %s, minimum value exceeded %d", val, v.InRange.Min)
	}

	return nil
}

type RegexValidation struct {
	Regex regexp.Regexp
}

func (v *RegexValidation) Validate(val string) error {
	if !v.Regex.MatchString(val) {
		return fmt.Errorf("Value %s, doesn't match with %s regex", val, v.Regex.String())
	}

	return nil
}
