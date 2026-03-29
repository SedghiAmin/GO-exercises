package luhn

import(
    "strings"
    "strconv"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) < 2 {
		return false
	}
	j := 1
	sum := 0
	for i := len(id) - 1; i >= 0; i-- {
		charStr := string(id[i])
		v, err := strconv.Atoi(charStr)
		if err != nil {
			return false
		}
		if j%2 == 0 {
			if v *= 2; v > 9 {
				v -= 9
			}
		}
		sum += v
		j++
	}
	if sum%10 == 0 {
		return true
	}
	return false
}
