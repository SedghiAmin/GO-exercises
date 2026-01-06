package thefarm

import (
    "fmt"
    "errors"
    )

// TODO: define the 'DivideFood' function
type InvalidCowsError struct {
	cows int
	msg  string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.msg)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows: cows, msg: "there are no negative cows"}
	} else if cows == 0 {
		return &InvalidCowsError{cows: cows, msg: "no cows don't need food"}
	}
	return nil
}

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	totalAmount, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return totalAmount * factor / float64(cows), nil

}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if err := ValidateNumberOfCows(cows); err != nil {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cows)
}
