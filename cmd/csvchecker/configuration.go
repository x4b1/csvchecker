package main

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/xabi93/csvchecker"
)

const (
	stringValidator = "string"
	numberValidator = "number"
	regexpValidator = "regexp"
)

type configuration struct {
	Separator  string `json:"separator"`
	WithHeader bool   `json:"withHeader"`
	Columns    []struct {
		Position  int                     `json:"position"`
		Validator *validatorConfiguration `json:"validation"`
	} `json:"columns"`
}

func (conf *configuration) createChecker() *csvchecker.Checker {
	checker := csvchecker.NewChecker(rune(conf.Separator[0]), conf.WithHeader)
	for _, col := range conf.Columns {
		checker.AddColum(csvchecker.NewColumn(col.Position, col.Validator.createValidator()))
	}

	return checker
}

func (conf *configuration) validate() error {
	if len(conf.Separator) != 1 {
		return errors.New("Invalid 'separator'")
	}
	for i, col := range conf.Columns {
		if col.Position < 1 {
			return fmt.Errorf("Invalid 'position' in column %d", i+1)
		}
		if !col.Validator.existValidator() {
			return fmt.Errorf("validation '%s' does not exists", col.Validator.ValidatorType)
		}
	}
	return nil
}

type validatorConfiguration struct {
	ValidatorType string `json:"type"`
	AllowEmpty    bool   `json:"allowEmpty"`
	CheckRange    *struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"range"`
	Regex string
}

func (c *validatorConfiguration) existValidator() bool {
	for _, a := range []string{stringValidator, numberValidator, regexpValidator} {
		if a == c.ValidatorType {
			return true
		}
	}
	return false
}

func (c *validatorConfiguration) createValidator() csvchecker.Validator {
	switch c.ValidatorType {
	case stringValidator:
		if c.CheckRange != nil {
			return csvchecker.NewStringValidation(c.AllowEmpty, csvchecker.NewRangeValidation(c.CheckRange.Min, c.CheckRange.Max))
		}
		return csvchecker.NewStringValidation(c.AllowEmpty, nil)
	case numberValidator:
		if c.CheckRange != nil {
			return csvchecker.NewNumberValidation(c.AllowEmpty, csvchecker.NewRangeValidation(c.CheckRange.Min, c.CheckRange.Max))
		}
		return csvchecker.NewNumberValidation(c.AllowEmpty, nil)
	case regexpValidator:
		return csvchecker.NewRegexpValidation(regexp.MustCompile(c.Regex))
	}

	return nil
}
