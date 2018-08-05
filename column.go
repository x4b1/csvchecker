package csvchecker

type column struct {
	position  int
	validator Validator
}

func NewColumn(position int, validator Validator) *column {
	return &column{position, validator}
}
