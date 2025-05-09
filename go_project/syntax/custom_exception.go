package syntax

type DivisionByZero struct{}

func (e DivisionByZero) Error() string {
	return "Can't Divide By Zero"
}
