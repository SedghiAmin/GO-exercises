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
	result := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			result = result.Append(IntList{v})
		}
	}
	return result
}

func (s IntList) Length() int {
	var l int = 0
	for range s {
		l++
	}
	return l
}

func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, 0)
	for _, v := range s {
		result = result.Append(IntList{fn(v)})
	}
	return result
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
	totalLen := 0
	for _, lst := range lists {
		totalLen += len(lst)
	}

	result := make(IntList, 0, totalLen)
	result = s
	for _, lst := range lists {
		result = result.Append(lst)
	}
	return result
}

func main() {
	var source IntList = []int{1, 2, 3}
	var ext IntList = []int{4, 5, 6}
	var lst []IntList = []IntList{
		{5, 6, 3},
		{9, 4, 1},
	}
	fn := func(x int) bool {
		if x%2 == 0 {
			return true
		}
		return false
	}
	fmt.Println(source.Append(ext)) //[1 2 3 4 5 6]
	fmt.Println(source.Concat(lst)) //[1 2 3 5 6 3 9 4 1]
	fmt.Println(source.Filter(fn))

	fn2 := func(x int) int {
		return x * x
	}
	fmt.Println(source.Map(fn2))
}
