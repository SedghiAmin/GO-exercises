package main

import (
	"errors"
	"fmt"
)

var (
	ErrStop        = errors.New("stop codon")
	ErrInvalidBase = errors.New("invalid codon")
)

var acid map[string]string = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}
	r := []string{}
	str := ""
	for i, char := range rna {
		str += string(char)
		if (i+1)%3 == 0 {
			protein, err := FromCodon(str)
			if err != nil {
				if err == ErrStop {
					return r, nil
				}
				return r, err
			}
			r = append(r, protein)
			str = ""
		}
	}
	return r, nil
}
