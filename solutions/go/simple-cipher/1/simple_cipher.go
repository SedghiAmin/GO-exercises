package cipher

import (
	"strings"
	"unicode"
)

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

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}
	
	allA := true
	for _, chr := range key {
		if chr < 'a' || chr > 'z' {
			return nil
		}
		if chr != 'a' {
			allA = false
		}
	}
	
	if allA {
		return nil
	}
	
	return vigenere(key)
}

func (v vigenere) Encode(input string) string {
	var output strings.Builder
	keyLen := len(v)
	keyIndex := 0
	
	for _, chr := range input {
		if unicode.IsLetter(chr) {
			lower := unicode.ToLower(chr)
			shift := int(v[keyIndex%keyLen] - 'a')
			
			encoded := (int(lower-'a') + shift) % 26
			output.WriteRune(rune(encoded) + 'a')
			
			keyIndex++
		}
	}
	return output.String()
}

func (v vigenere) Decode(input string) string {
	var output strings.Builder
	keyLen := len(v)
	keyIndex := 0
	
	for _, chr := range input {
		if unicode.IsLetter(chr) {
			lower := unicode.ToLower(chr)
			shift := int(v[keyIndex%keyLen] - 'a')
			
			decoded := (int(lower-'a') - shift) % 26
			if decoded < 0 {
				decoded += 26
			}
			
			output.WriteRune(rune(decoded) + 'a')
			keyIndex++
		}
	}
	return output.String()
}