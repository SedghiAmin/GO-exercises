package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	rnd            *rand.Rand
	suits          []string
	drawnCards     map[string]bool
	cardsPerDeck   int
	numberOfDecks  int
	totalCards     int
	reshuffleAt    int
	reshuffleCount int
}

func NewGame(suits []string, cardsPerDeck int, numberOfDecks int) *Game {
	source := rand.NewSource(time.Now().UnixNano())

	return &Game{
		rnd:            rand.New(source),
		suits:          suits,
		drawnCards:     make(map[string]bool),
		cardsPerDeck:   cardsPerDeck,
		numberOfDecks:  numberOfDecks,
		totalCards:     cardsPerDeck * numberOfDecks,
		reshuffleAt:    cardsPerDeck * numberOfDecks * 80 / 100,
		reshuffleCount: 0,
	}
}

func FetchCard(g *Game) string {
	if len(g.drawnCards) >= g.reshuffleAt {
		g.Reshuffle()
	}
	var fetch string = ""
	for {
		cardNum := g.rnd.Intn(13) + 1 // 1-13
		suitIndex := g.rnd.Intn(4)    // 0-3
		deck := g.rnd.Intn(g.numberOfDecks) + 1

		key := fmt.Sprintf("%d-%d-%s", deck, cardNum, g.suits[suitIndex])

		if !g.drawnCards[key] {
			g.drawnCards[key] = true
			fetch = fmt.Sprintf("Deck%d-%s%d,", deck, g.suits[suitIndex], cardNum)
			break
		}
	}
	return fetch
}

func (g *Game) Reshuffle() {
	g.reshuffleCount++

	g.drawnCards = make(map[string]bool)

	source := rand.NewSource(time.Now().UnixNano())

	g.rnd = rand.New(source)

}

func main() {

	cardsPerDeck := 52
	numberOfDecks := 1
	maxOfPlayer := 5
	numberOfPlayers := rand.Intn(maxOfPlayer) + 1

	suits := []string{"Heart", "Diamond", "Spades", "Clubs"}

	var gameSet = NewGame(suits, cardsPerDeck, numberOfDecks)

	i := 0
	for {
		if numberOfPlayers > 2 {
			i++

			fmt.Printf("Set %d - number of player(s): %d\nDelear Cards: %s ** %s \n", i, numberOfPlayers-1, FetchCard(gameSet), FetchCard(gameSet))

			for j := 1; j < numberOfPlayers; j++ {
				fmt.Printf("   Player %d: %s ** %s \n", j, FetchCard(gameSet), FetchCard(gameSet))
			}
			fmt.Printf("\n\n*********************************\n\n")
		} else {
			fmt.Println("There are no player(s) ... !!")
			fmt.Printf("Number of Reshuffle: %d\n", gameSet.reshuffleCount)
			fmt.Printf("Number of Set(s): %d", i)
			fmt.Printf("\n\n*********************************\n\n")
			break
		}
		numberOfPlayers = rand.Intn(maxOfPlayer) + 1
	}
}
