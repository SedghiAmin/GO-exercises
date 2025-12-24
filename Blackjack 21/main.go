package main

import (
	"fmt"
	"math/rand"
	"time"
)

func FetchCard(rng *rand.Rand, suits []string, drawnCards map[string]bool) string {
	l := true
	var fetch string = ""
	for l {
		cardNum := rng.Intn(13) + 1 // 1-13
		suitIndex := rng.Intn(4)    // 0-3
		deck := rng.Intn(6) + 1

		key := fmt.Sprintf("%d-%d-%s", deck, cardNum, suits[suitIndex])

		if !drawnCards[key] {
			drawnCards[key] = true
			fetch = fmt.Sprintf("Deck %d: %s %d", deck, suits[suitIndex], cardNum)
			l = false
		}
	}
	return fetch
}

func main() {
	// create generator
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	cardsPerDeck := 52
	numberOfDecks := 6

	suits := []string{"Heart", "Diamond", "Spades", "Clubs"}
	drawnCards := make(map[string]bool)

	fmt.Println("Random Cards: ")
	for i := 0; i < cardsPerDeck*numberOfDecks; i++ {
		fmt.Printf("%d: %s\n", i+1, FetchCard(rng, suits, drawnCards))
	}
}
