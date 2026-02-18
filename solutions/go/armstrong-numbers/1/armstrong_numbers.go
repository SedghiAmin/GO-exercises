package armstrong
import "math"
func IsNumber(n int) bool {
	if n == 0 {
		return true
	}
	if n < 0 {
		n = -n
	}
	digits := make([]int, 0)
	number := n
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	sum := 0.0
	for i := range digits {
		sum += math.Pow(float64(digits[i]), float64(len(digits)))
	}
	if sum == float64(number) {
		return true
	}
	return false
}
