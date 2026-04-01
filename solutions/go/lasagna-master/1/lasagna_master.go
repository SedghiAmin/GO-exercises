package lasagna

func PreparationTime(layers []string, prepartionTime int) int {
	if prepartionTime > 0 {
		return len(layers) * prepartionTime
	}
	return len(layers) * 2
}

func Quantities(layers []string) (int, float64) {
	var s, n int
	for _, l := range layers {
		if l == "sauce" {
			s += 1
		} else if l == "noodles" {
			n += 1
		}
	}
	return n * 50, float64(s) * 0.2
}

func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(twoPortions []float64, portionsCount int) []float64 {
	newPortions := make([]float64, 0, len(twoPortions))
	for _, portion := range twoPortions {
		newPortions = append(newPortions, float64(portionsCount)/2*portion)
	}
	return newPortions
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
