package main

import "fmt"

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	panic("Please implement the Foldl function")
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	panic("Please implement the Foldr function")
}

func (s IntList) Filter(fn func(int) bool) IntList {
	panic("Please implement the Filter function")
}

func (s IntList) Length() int {
	panic("Please implement the Length function")
}

func (s IntList) Map(fn func(int) int) IntList {
	panic("Please implement the Map function")
}

func (s IntList) Reverse() IntList {
	panic("Please implement the Reverse function")
}

func (s IntList) Append(lst IntList) IntList {
	out := make(IntList, len(s)+len(lst))
	for i := 0; i < len(s); i++ {
		out[i] = s[i]
	}
	for i := 0; i < len(lst); i++ {
		out[len(s)+i] = lst[i]
	}
	return out
}

func (s IntList) Concat(lists []IntList) IntList {
	panic("Please implement the Concat function")
}

func main() {
	var source IntList = []int{1, 2, 3}
	var lst IntList = []int{4, 5, 6}
	fmt.Println(source.Append(lst))
}
