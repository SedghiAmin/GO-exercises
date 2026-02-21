package prime

func Factors(n int64) []int64 {
	out := make([]int64, 0)
	var i int64
	for i = 2; i <= n; i++ {
		for n%i == 0 {
			out = append(out, i)
			n = n / i
		}
	}
	return out
}
