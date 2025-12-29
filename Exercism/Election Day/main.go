package main

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}
