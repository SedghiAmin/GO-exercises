package main

import (
	"fmt"
	"math/rand"
	"time"
)

type GameSet struct {
	rng           *rand.Rand
	suits         []string
	drawnCards    map[string]bool
	cardsPerDeck  int
	numberOfDecks int
	totalCards    int
	reshuffleAt   int
}

func NewGameSet(suits []string, drawnCards map[string]bool, cardsPerDeck int, numberOfDecks int) *GameSet {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	return &GameSet{
		rng:           rng,
		suits:         suits,
		drawnCards:    drawnCards,
		cardsPerDeck:  cardsPerDeck,
		numberOfDecks: numberOfDecks,
		totalCards:    cardsPerDeck * numberOfDecks,
		reshuffleAt:   cardsPerDeck * numberOfDecks * 80 / 100,
	}
}

func FetchCard(rng *rand.Rand, suits []string, drawnCards map[string]bool) string {
	var fetch string = ""
	for {
		cardNum := rng.Intn(13) + 1 // 1-13
		suitIndex := rng.Intn(4)    // 0-3
		deck := rng.Intn(6) + 1

		key := fmt.Sprintf("%d-%d-%s", deck, cardNum, suits[suitIndex])

		if !drawnCards[key] {
			drawnCards[key] = true
			fetch = fmt.Sprintf("Deck %d: %s %d", deck, suits[suitIndex], cardNum)
			break
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
