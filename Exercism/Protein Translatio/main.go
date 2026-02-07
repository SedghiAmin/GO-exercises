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

func FromCodon(codon string) (string, error) {
	if protein, ok := acid[codon]; ok {
		if protein == "STOP" {
			return "", ErrStop
		}
		return protein, nil
	}
	return "", ErrInvalidBase
}

func main() {

	fmt.Println("=== Basic Tests ===")

	rna1 := "AUGUUUUCU"
	proteins1, err1 := FromRNA(rna1)
	fmt.Printf("RNA: %s\n", rna1)
	if err1 != nil {
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("Proteins: %v\n", proteins1)
	}

	fmt.Println()

	rna2 := "AUGUUUUAAUCU"
	proteins2, err2 := FromRNA(rna2)
	fmt.Printf("RNA: %s\n", rna2)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	} else {
		fmt.Printf("Proteins: %v\n", proteins2)
	}

	fmt.Println()

	rna3 := "AUGXYZUCU"
	proteins3, err3 := FromRNA(rna3)
	fmt.Printf("RNA: %s\n", rna3)
	if err3 != nil {
		fmt.Printf("Error: %v\n", err3)
	} else {
		fmt.Printf("Proteins: %v\n", proteins3)
	}
}
