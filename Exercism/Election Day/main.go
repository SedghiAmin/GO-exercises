package main

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

func main() {
	var initialVotes int
	initialVotes = 2

	var counter *int
	counter = NewVoteCounter(initialVotes)
	fmt.Println(*counter == initialVotes) // true

	var votes int
	votes = 3

	var voteCounter *int
	voteCounter = &votes

	fmt.Println(VoteCount(voteCounter))
	// => 3

	var nilVoteCounter *int
	fmt.Println(VoteCount(nilVoteCounter))
	// => 0

	IncrementVoteCount(voteCounter, 2)

	fmt.Println(votes == 5)        // true
	fmt.Println(*voteCounter == 5) // true
}
