package main

import(
	"fmt"
)

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](slc []T, x T) int{
	for i, val:= range slc{
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if val == x{
			return i
		}
	}
	return -1
}

func main(){
	// Index works on a slice of ints
	slc1:= []int{10,13,17,23}
	fmt.Println(Index(slc1, 13))

	// Index also works on a slice of strings
	slc2 := []string{"foo", "bar", "baz"}
	fmt.Println(Index(slc2, "hello"))
}