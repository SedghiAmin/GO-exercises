
package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
    remark = strings.TrimSpace(remark)
    input := []rune(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}else if remark == strings.ToUpper(remark) && input[len(input)-1] == '?' && haveletter(remark) {
		return "Calm down, I know what I'm doing!"
	} else if remark == strings.ToUpper(remark) && haveletter(remark) {
		return "Whoa, chill out!"
	}else if input[len(input)-1] == '?' {
		return "Sure."
	}
	return "Whatever."
}
func haveletter(remark string) bool{
    input := []rune(remark)
    for _, char:= range input{
       if unicode.IsLetter(char){
           return true
        }
    }
    return false
}