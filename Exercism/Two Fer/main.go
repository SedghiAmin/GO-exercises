package main

import "fmt"

func ShareWith(name string) string {
	var nam string
	if name != "" {
		nam = name
	} else {
		nam = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", nam)
}

func main() {
	fmt.Println(ShareWith(""))
}
