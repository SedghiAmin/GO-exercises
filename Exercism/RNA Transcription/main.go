package main

import "fmt"

func ToRNA(dna string) string {
	rna := make([]byte, len(dna))
	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'G':
			rna[i] = 'C'
		case 'C':
			rna[i] = 'G'
		case 'T':
			rna[i] = 'A'
		case 'A':
			rna[i] = 'U'
		default:
			rna[i] = dna[i]
		}
	}
	return string(rna)
}

func main() {

	dna1 := "GCTA"
	rna1 := ToRNA(dna1)
	fmt.Printf("DNA: %s -> RNA: %s\n", dna1, rna1)
	// output: DNA: GCTA -> RNA: CGAU

	dna2 := "ACGTGGTCTTAA"
	rna2 := ToRNA(dna2)
	fmt.Printf("DNA: %s -> RNA: %s\n", dna2, rna2)
	// output: DNA: ACGTGGTCTTAA -> RNA: UGCACCAGAAUU

	dna3 := "GCTAX"
	rna3 := ToRNA(dna3)
	fmt.Printf("DNA: %s -> RNA: %s\n", dna3, rna3)
	// output: DNA: GCTAX -> RNA: CGAUX
}
