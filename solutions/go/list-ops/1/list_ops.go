package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := s.Length() - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
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
	result := make(IntList, 0)
	for i := len(s) - 1; i >= 0; i-- {
		result = result.Append(IntList{s[i]})
	}
	return result
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