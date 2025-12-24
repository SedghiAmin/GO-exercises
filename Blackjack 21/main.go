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

func FetchCard(g *GameSet) string {
	var fetch string = ""
	for {
		cardNum := g.rng.Intn(13) + 1 // 1-13
		suitIndex := g.rng.Intn(4)    // 0-3
		deck := g.rng.Intn(6) + 1

		key := fmt.Sprintf("%d-%d-%s", deck, cardNum, g.suits[suitIndex])

		if !g.drawnCards[key] {
			g.drawnCards[key] = true
			fetch = fmt.Sprintf("Deck %d: %s %d", deck, g.suits[suitIndex], cardNum)
			break
		}
	}
	return fetch
}

func main() {

	cardsPerDeck := 52
	numberOfDecks := 6

	suits := []string{"Heart", "Diamond", "Spades", "Clubs"}
	drawnCards := make(map[string]bool)

	var gameSet = NewGameSet(suits, drawnCards, cardsPerDeck, numberOfDecks)

	fmt.Println("Random Cards: ")
	for i := 0; i < cardsPerDeck*numberOfDecks; i++ {
		fmt.Printf("%d: %s\n", i+1, FetchCard(gameSet))
	}
}
