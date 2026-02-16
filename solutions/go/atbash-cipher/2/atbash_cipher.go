package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var output strings.Builder
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)
	for _, c := range s {
		if unicode.IsLetter(c) {
			offset := c - 'a'
			output.WriteRune('z' - offset)
		} else if unicode.IsDigit(c) {
			output.WriteRune(c)
		}
	}
    if output.String() == ""{
        return ""
    }
    var result strings.Builder
    for i, c:= range output.String(){
        if i > 0 && i % 5 == 0{
            result.WriteRune(' ')
        }
        result.WriteRune(c)
    }
	return result.String()
}