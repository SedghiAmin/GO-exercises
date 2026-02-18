package main

import (
	"errors"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	runes := []rune(phoneNumber)
	output := make([]rune, 0, len(runes))
	for i, ch := range runes {
		if i == 0 && ch == '+' {
			continue
		}
		if i == 1 && runes[0] == '+' {
			if ch != '1' {
				return "", errors.New("Invalid country code")
			}
		}
		if unicode.IsDigit(ch) {
			output = append(output, ch)
		} else if ch == '(' || ch == ')' || ch == '-' || ch == '.' || ch == ' ' {
			continue
		} else {
			return "", errors.New("invalid with punctuations")
		}
	}
	if len(output) == 11 {
		if output[0] == '1' {
			output = output[1:]
		} else {
			return "", errors.New("invalid when 11 digits does not start with a 1")
		}
	}
	if len(output) != 10 {
		return "", errors.New("invalid count of digits digits")
	}
	if int(output[0]-'0') < 2 {
		return "", errors.New("invalid area code")
	}
	if int(output[3]-'0') < 2 {
		return "", errors.New("invalid exchange code")
	}
	return string(output), nil
}

func AreaCode(phoneNumber string) (string, error) {
	value, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	output := []rune(value)
	output = output[:3]
	return string(output), nil
}

func Format(phoneNumber string) (string, error) {
	value, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	runes := []rune(value)
	output := make([]rune, 0)
	output = append(output, '(')
	output = append(output, runes[:3]...)
	output = append(output, ')')
	output = append(output, ' ')
	output = append(output, runes[3:6]...)
	output = append(output, '-')
	output = append(output, runes[6:]...)
	return string(output), nil
}
