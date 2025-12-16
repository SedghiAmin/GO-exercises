package main

import "fmt"

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
	myList[len(myList)-1] = friendList[len(myList)-1]
}

func ScaleRecipe(twoPortions []float64, portionsCount int) []float64 {
	newPortions := make([]float64, 0, len(twoPortions))
	for _, portion := range twoPortions {
		newPortions = append(newPortions, float64(portionsCount)/2*portion)
	}
	return newPortions
}

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println(PreparationTime(layers, 3))
	// => 18
	fmt.Println(PreparationTime(layers, 0))
	// => 12
	fmt.Println(Quantities(layers))
	// => 100, 0.4
	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella"}
	AddSecretIngredient(friendsList, myList)
	fmt.Printf("%#v\n", myList)
	// myList => []string{"noodles", "meat", "sauce", "mozzarella", "kampot pepper"})
	quantities := []float64{1.2, 3.6, 10.5}
	fmt.Printf("%#v\n", ScaleRecipe(quantities, 4))
	// => []float64{ 2.4, 7.2, 21 }
}
