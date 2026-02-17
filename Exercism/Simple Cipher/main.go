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

func main() {
	fmt.Println("=== Testing Caesar Cipher (shift 3) ===")
	caesar := NewCaesar()

	testCases := []string{
		"hello world",
		"HELLO WORLD",
		"Hello, World! 123",
	}

	for _, tc := range testCases {
		encoded := caesar.Encode(tc)
		decoded := caesar.Decode(encoded)
		fmt.Printf("Original: %q\n", tc)
		fmt.Printf("Encoded:  %q\n", encoded)
		fmt.Printf("Decoded:  %q\n", decoded)
		fmt.Printf("Match:    %v\n\n", decoded == "helloworld")
	}

	fmt.Println("=== Testing Shift Cipher (shift 5) ===")
	shift5 := NewShift(5)
	if shift5 != nil {
		text := "the quick brown fox"
		encoded := shift5.Encode(text)
		decoded := shift5.Decode(encoded)
		fmt.Printf("Original: %q\n", text)
		fmt.Printf("Encoded:  %q\n", encoded)
		fmt.Printf("Decoded:  %q\n", decoded)
		fmt.Printf("Match:    %v\n\n", decoded == text)
	}

	fmt.Println("=== Testing Shift Cipher (shift -3) ===")
	shiftMinus3 := NewShift(-3)
	if shiftMinus3 != nil {
		text := "the quick brown fox"
		encoded := shiftMinus3.Encode(text)
		decoded := shiftMinus3.Decode(encoded)
		fmt.Printf("Original: %q\n", text)
		fmt.Printf("Encoded:  %q\n", encoded)
		fmt.Printf("Decoded:  %q\n", decoded)
		fmt.Printf("Match:    %v\n\n", decoded == text)
	}

	fmt.Println("=== Testing Vigenere Cipher (key: 'abc') ===")
	vigenere1 := NewVigenere("abc")
	if vigenere1 != nil {
		text := "hello world"
		encoded := vigenere1.Encode(text)
		decoded := vigenere1.Decode(encoded)
		fmt.Printf("Original: %q\n", text)
		fmt.Printf("Key:      %q\n", "abc")
		fmt.Printf("Encoded:  %q\n", encoded)
		fmt.Printf("Decoded:  %q\n", decoded)
		fmt.Printf("Match:    %v\n\n", decoded == "helloworld")
	}

	fmt.Println("=== Testing Vigenere Cipher (key: 'gold') ===")
	vigenere2 := NewVigenere("gold")
	if vigenere2 != nil {
		text := "attack at dawn"
		encoded := vigenere2.Encode(text)
		decoded := vigenere2.Decode(encoded)
		fmt.Printf("Original: %q\n", text)
		fmt.Printf("Key:      %q\n", "gold")
		fmt.Printf("Encoded:  %q\n", encoded)
		fmt.Printf("Decoded:  %q\n", decoded)
		fmt.Printf("Match:    %v\n\n", decoded == "attackatdawn")
	}

	fmt.Println("=== Testing Invalid Inputs ===")
	fmt.Printf("NewShift(0):   %v\n", NewShift(0))
	fmt.Printf("NewShift(26):  %v\n", NewShift(26))
	fmt.Printf("NewShift(-26): %v\n", NewShift(-26))
	fmt.Printf("NewVigenere(\"\"):     %v\n", NewVigenere(""))
	fmt.Printf("NewVigenere(\"ABC\"):   %v\n", NewVigenere("ABC"))
	fmt.Printf("NewVigenere(\"aaa\"):   %v\n", NewVigenere("aaa"))
	fmt.Printf("NewVigenere(\"123\"):   %v\n", NewVigenere("123"))
}
