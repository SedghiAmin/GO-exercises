package etl
import "strings"
func Transform(in map[int][]string) map[string]int {
	out:= make(map[string]int)
    for score, str:= range in{
        for _, char:= range str{
            lChar:= strings.ToLower(char)
            out[lChar] = score
        }
    }
    return out
}
