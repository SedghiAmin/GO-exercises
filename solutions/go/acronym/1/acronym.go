package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	out := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if i == 0 && string(s[i]) != " " && string(s[i]) != "-" && string(s[i]) != "_" {
			out.WriteString(string(s[i]))
			continue
		}
		if string(s[i]) == " " || string(s[i]) == "-" || string(s[i]) == "_" {
			if (i+1 < len(s)) && (string(s[i+1]) != " ") && (string(s[i+1]) != "-") && (string(s[i+1]) != "_") {
				out.WriteString(strings.ToUpper(string(s[i+1])))
			}
		}
	}
	return out.String()
}
