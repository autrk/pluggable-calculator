package operation

import "errors"

func Divide(a, b float64) (float64, error){
	if b==0 {
		return 0, errors.New("Division by zero is not permitted")
	}
	return a / b, nil
}
