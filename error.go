package csvchecker

import "fmt"

type CSVError interface {
	ToString() string
}

type rowError struct {
	err error
}

func (r *rowError) ToString() string {
	return r.err.Error()
}

type colError struct {
	line int
	col  int
	err  error
}

func (r *colError) ToString() string {
	return fmt.Sprintf("Error on row %d and column %d: %s", r.line, r.col, r.err.Error())
}
