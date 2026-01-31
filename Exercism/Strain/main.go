package main

func Keep[T any](arr []T, f func(T) bool) []T {
	var r []T

	for _, v := range arr {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}
