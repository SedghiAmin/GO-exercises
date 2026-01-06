package lasagna

import("fmt")

// TODO: define the 'OvenTime' constant
const OvenTime = 40
// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    r:= numberOfLayers * 2
	return r
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return numberOfLayers * 2 + actualMinutesInOven
}

func lasagna(){
	fmt.Println(RemainingOvenTime(15))
    fmt.Println(PreparationTime(5))
    fmt.Println(ElapsedTime(5,20))
}
