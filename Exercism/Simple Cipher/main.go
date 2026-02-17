package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type shift int
type vigenere string

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance > 25 || distance == 0 {
		return nil
	}
	return shift(distance)
}
