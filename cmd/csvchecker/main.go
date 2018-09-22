package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/xabi93/csvchecker"
)

const (
	stringValidator = "string"
	numberValidator = "number"
	regexpValidator = "regexp"
)

var usage = func() {
	fmt.Fprint(os.Stderr, `usage: csvchecker (configuration JSON path) (CSV path)
csvchecker given a configuration json and csv validate structure`)
}

func main() {
	h := commandHandler{}
	h.fromArgs(os.Args[1:])
}

type commandHandler struct {
}

func (c *commandHandler) fromArgs(args []string) {
	p := c.getArgs(args)
	if len(p) < 2 {
		c.fail("Invalid argument number")
	}
	conf := c.processConfigurationFile(p[0])

	doctor := conf.createDoctor()

	errors := doctor.Check(c.getFileReader(p[1]))
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Fprint(os.Stderr, err.ToString(), "\n")
		}
		c.fail("CSV file has some errors")
	}

}

func (c *commandHandler) getArgs(args []string) []string {
	flagSet := flag.NewFlagSet("csvchecker", flag.ExitOnError)
	flagSet.Usage = usage
	flagSet.Parse(args)

	return flagSet.Args()
}

func (c *commandHandler) fail(text string) {
	fmt.Fprint(os.Stderr, text)
	os.Exit(-1)
}

func (c *commandHandler) processConfigurationFile(path string) *configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		c.fail(err.Error())
	}

	conf := new(configuration)
	err = json.Unmarshal(file, &conf)

	if err != nil {
		c.fail(err.Error())
	}

	err = conf.validate()

	if err != nil {
		c.fail(err.Error())
	}

	return conf
}

func (c *commandHandler) getFileReader(path string) io.Reader {
	var r io.Reader
	var err error
	r, err = os.Open(path)

	if err != nil {
		c.fail(err.Error())
	}

	return r
}

type configuration struct {
	Separator  string `json:"separator"`
	WithHeader bool   `json:"withHeader"`
	Columns    []struct {
		Position  int                     `json:"position"`
		Validator *validatorConfiguration `json:"validation"`
	} `json:"columns"`
}

func (conf *configuration) createDoctor() *csvchecker.Checker {
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
		v := csvchecker.StringValidation{AllowEmpty: c.AllowEmpty}
		if c.CheckRange != nil {
			v.InRange = &csvchecker.RangeValidation{c.CheckRange.Min, c.CheckRange.Max}
		}
		return &v
	case numberValidator:
		v := csvchecker.NumberValidation{AllowEmpty: c.AllowEmpty}
		if c.CheckRange != nil {
			v.InRange = &csvchecker.RangeValidation{c.CheckRange.Min, c.CheckRange.Max}
		}
		return &v
	case regexpValidator:
		v := csvchecker.RegexValidation{*regexp.MustCompile(c.Regex)}
		return &v
	}

	return nil
}
