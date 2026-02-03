package main

import "fmt"

func RotationalCipher(plain string, shiftKey int) string {
	result := make([]rune, len(plain))
	for i, char := range plain {
		if plain[i] >= 'A' && plain[i] <= 'Z' {
			result[i] = 'A' + (char-'A'+rune(shiftKey))%26
		} else if plain[i] >= 'a' && plain[i] <= 'z' {
			result[i] = 'a' + (char-'a'+rune(shiftKey))%26
		} else {
			result[i] = char
		}
	}
	return string(result)
}

func main() {
	tests := []struct {
		input  string
		key    int
		output string
	}{
		{"omg", 5, "trl"},
		{"c", 0, "c"},
		{"Cool", 26, "Cool"},
		{"The quick brown fox jumps over the lazy dog.", 13,
			"Gur dhvpx oebja sbk whzcf bire gur ynml qbt."},
		{"Gur dhvpx oebja sbk whzcf bire gur ynml qbt.", 13,
			"The quick brown fox jumps over the lazy dog."},
		{"OMG", 5, "TRL"},
		{"Testing 1 2 3 testing!", 4, "Xiwxmrk 1 2 3 xiwxmrk!"},
		{"Let's eat, Grandma!", 21, "Gzo'n zvo, Bmviyhz!"},
	}

	for _, test := range tests {
		result := RotationalCipher(test.input, test.key)
		status := "✓"
		if result != test.output {
			status = "✗"
		}
		fmt.Printf("%s ROT%d %-45s → %-45s (Expected: %s)\n",
			status, test.key, test.input, result, test.output)
	}
}
