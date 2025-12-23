package main

import (
	"fmt"
)

type InvalidCowsError struct {
	cows int
	msg  string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.msg)
}

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
	if cows <= 0 {
		return 0, fmt.Errorf("invalid number of cows")
	}
	return DivideFood(fc, cows)
}

func main() {
	food := Food{}
	//fmt.Println(DivideFood(food, 5))
	fmt.Println(ValidateInputAndDivideFood(food, 5))
}
