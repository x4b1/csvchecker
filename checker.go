package csvchecker

import (
	"encoding/csv"
	"io"
)

type checker struct {
	separator  rune
	withHeader bool
	columns    []*column
}

func NewChecker(separator rune, withHeader bool) *checker {
	return &checker{
		separator:  separator,
		withHeader: withHeader,
	}
}

func (c *checker) AddColum(col *column) *checker {
	c.columns = append(c.columns, col)

	return c
}

func (c *checker) Check(reader io.Reader) []csvError {
	var errors []csvError
	lineNum := 0
	r := c.getReader(reader)

	for {
		lineNum++
		line, err := r.Read()

		if lineNum == 1 {
			continue
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			errors = append(errors, &rowError{
				err,
			})
		}
		errors = append(errors, c.checkLine(line, lineNum)...)
	}

	return errors
}

func (c *checker) checkLine(l []string, lNum int) []csvError {
	var errors []csvError
	for _, col := range c.columns {
		err := col.validator.Validate(l[col.position-1])
		if err != nil {
			errors = append(errors, &colError{
				lNum,
				col.position,
				err,
			})
		}
	}

	return errors
}

func (c *checker) getReader(reader io.Reader) *csv.Reader {
	r := csv.NewReader(reader)
	r.Comma = c.separator
	r.LazyQuotes = true
	return r
}
