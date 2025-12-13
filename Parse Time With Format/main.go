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
		"MM":   "04",
		"mm":   "01",
		"dd":   "02",
		"SS":   "05",
	}
	for k, v := range f {
		format = strings.ReplaceAll(format, k, v)
	}
	return time.Parse(format, strDate)
}

func main() {
	t, _ := ParseWithFormat("Tue, 09/22/1995, 13:00", "ddd, mm/dd/yyyy, HH:MM")
	fmt.Println(t)
}
