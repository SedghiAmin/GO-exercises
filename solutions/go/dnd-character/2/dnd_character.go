package dndcharacter

import (
    "math"
	"math/rand"
	"time"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score - 10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	
	sum := 0
	dice := make([]int, 4)

	dice[0] = rand.Intn(6) + 1
	for j := 1; j < 4; j++ {
		dice[j] = rand.Intn(5) + 1
		if dice[j] < dice[0] {
			tmp := dice[j]
			dice[j] = dice[0]
			dice[0] = tmp
		}
		sum += dice[j]
	}

	return sum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	rand.Seed(time.Now().UnixNano())
	char := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	char.Hitpoints = 10 + Modifier(char.Constitution)
	return char
}