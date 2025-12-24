package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// create generator
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	cardsPerDeck := 52
	numberOfDecks := 6

	suits := []string{"Heart", "Diamond", "Spades", "Clubs"}
	drawnCards := make(map[string]bool)

	fmt.Println("Random Cards: ")
	for i := 0; i < cardsPerDeck*numberOfDecks; {
		cardNum := rng.Intn(13) + 1 // 1-13
		suitIndex := rng.Intn(4)    // 0-3
		deck := rng.Intn(6) + 1

		key := fmt.Sprintf("%d-%d-%s", deck, cardNum, suits[suitIndex])

		if !drawnCards[key] {
			drawnCards[key] = true
			fmt.Printf("Deck %d: %s %d\n", deck, suits[suitIndex], cardNum)
			i++
		}
	}
}
