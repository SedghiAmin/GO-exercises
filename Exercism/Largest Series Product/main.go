package main

import (
	"fmt"
	"slices"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span {
		return 0, fmt.Errorf("span must be smaller than string length")
	}
	if span < 0 {
		return 0, fmt.Errorf("span must not be negative")
	}
	var sum int64
	result := make([]int64, len(digits)-span+1)
	sum = 1
	digitsRunes := []rune(digits)
	for i := 0; i <= len(digitsRunes)-span; i++ {
		for j := i; j < i+span; j++ {
			if unicode.IsDigit(digitsRunes[j]) {
				sum *= int64(digitsRunes[j] - '0')
			} else {
				return 0, fmt.Errorf("digits input must only contain digits")
			}
		}
		result[i] = sum
		sum = 1
	}
	return slices.Max(result), nil
}

func main() {
	fmt.Println(LargestSeriesProduct("12345", 5))                                              //0 span must be smaller than string length
	fmt.Println(LargestSeriesProduct("12345", -1))                                             //0 span must not be negative
	fmt.Println(LargestSeriesProduct("1234a5", 2))                                             //0 digits input must only contain digits
	fmt.Println(LargestSeriesProduct("", 1))                                                   //0 span must be smaller than string length
	fmt.Println(LargestSeriesProduct("123", 4))                                                //0 span must be smaller than string length
	fmt.Println(LargestSeriesProduct("99099", 3))                                              // 0
	fmt.Println(LargestSeriesProduct("0000", 2))                                               // 0
	fmt.Println(LargestSeriesProduct("73167176531330624919225119674426574742355349194934", 6)) //23520 <nil>
	fmt.Println(LargestSeriesProduct("0123456789", 5))                                         //15120 <nil
	fmt.Println(LargestSeriesProduct("1027839564", 3))                                         //270 <nil
}
