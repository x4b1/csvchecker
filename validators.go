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
	min int
	max int
}

type StringValidation struct {
	allowEmpty bool
	inRange    *RangeValidation
}

func (v *StringValidation) Validate(val string) error {
	strLn := len(val)

	if !v.allowEmpty && strLn < 1 {
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

type NumberValidation struct {
	inRange *RangeValidation
}

func (v *NumberValidation) Validate(val string) error {
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

type RegexValidation struct {
	regex regexp.Regexp
}

func (v *RegexValidation) Validate(val string) error {
	if !v.regex.MatchString(val) {
		return fmt.Errorf("Value %s, doesn't match with %s regex", val, v.regex.String())
	}

	return nil
}
