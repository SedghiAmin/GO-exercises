package main

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
