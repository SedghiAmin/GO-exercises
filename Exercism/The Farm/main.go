package main

import (
	"fmt"
)

type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}

type Food struct{}

func (f Food) FodderAmount(amount int) (float64, error) {
	return 50.0, nil
}

func (f Food) FatteningFactor() (float64, error) {
	return 1.5, nil
}

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

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	err := ValidateNumberOfCows(cows)
	if err != nil {
		return 0, err
	}
	return DivideFood(fc, cows)
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

func main() {
	food := Food{}
	//fmt.Println(DivideFood(food, 5))
	v, e := ValidateInputAndDivideFood(food, 5)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(v)
	}
}
