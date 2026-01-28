package main

import "errors"

type Histogram map[byte]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for i := 0; i < len(d); i++ {
		switch d[i] {
		case 'A', 'C', 'G', 'T':
			h[d[i]]++
		default:
			return nil, errors.New(" invalid nucleotides")
		}
	}
	return h, nil
}
