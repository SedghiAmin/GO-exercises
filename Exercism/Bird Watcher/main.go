package main

func TotalBirdCount(birdsPerDay []int) int {
	var sum int
	for i := 0; i < len(birdsPerDay); i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	i := (week - 1) * 7
	var sum, n int
	if i+7 < len(birdsPerDay) {
		n = i + 7
	} else {
		n = len(birdsPerDay)
	}
	for j := i; j < n; j++ {
		sum += birdsPerDay[j]
	}
	return sum
}

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	println(TotalBirdCount(birdsPerDay))
	println(BirdsInWeek(birdsPerDay, 2))
}
