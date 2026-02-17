package main

import (
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

func (c shift) Encode(input string) string {
	var output strings.Builder
	shiftValue := int(c)

	for _, chr := range input {
		if unicode.IsLetter(chr) {
			lower := unicode.ToLower(chr)

			encoded := int(lower-'a') + shiftValue

			encoded %= 26
			if encoded < 0 {
				encoded += 26
			}

			output.WriteRune(rune(encoded) + 'a')
		}
	}
	return output.String()
}

func (c shift) Decode(input string) string {
	var output strings.Builder
	shiftValue := int(c)

	for _, chr := range input {
		if unicode.IsLetter(chr) {
			lower := unicode.ToLower(chr)

			decoded := int(lower-'a') - shiftValue

			decoded %= 26
			if decoded < 0 {
				decoded += 26
			}

			output.WriteRune(rune(decoded) + 'a')
		}
	}
	return output.String()
}
