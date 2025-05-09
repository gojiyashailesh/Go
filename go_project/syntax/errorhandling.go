package syntax

import "errors"

func ZeroDevisionError(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("devide By Zero")
	}
	return a / b, nil
}
