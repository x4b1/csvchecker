package csvchecker

type Column struct {
	position  int
	validator Validator
}

func NewColumn(position int, validator Validator) *Column {
	return &Column{position, validator}
}
