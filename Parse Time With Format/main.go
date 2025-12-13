package main

import (
	"fmt"
	"strings"
	"time"
)

func ParseWithFormat(strDate, format string) (time.Time, error) {

	f := map[string]string{
		"dddd": "Monday",
		"ddd":  "Mon",
		"yyyy": "2006",
		"yy":   "06",
		"HH":   "15",
		"MM":   "01",
		"mm":   "04",
		"dd":   "02",
		"SS":   "05",
	}
	for k, v := range f {
		format = strings.ReplaceAll(format, k, v)
	}
	return time.Parse(format, strDate)
}

func main() {
	t, _ := ParseWithFormat("Tue, 09/22/1995, 13:00", "ddd, MM/dd/yyyy, HH:mm")
	fmt.Println(t)
}
