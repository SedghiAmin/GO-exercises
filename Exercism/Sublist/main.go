package main

import "fmt"

// Relation is the comparison between lists
type Relation string

// Possible relations
const (
	RelationEqual     Relation = "equal"
	RelationSublist   Relation = "sublist"
	RelationSuperlist Relation = "superlist"
	RelationUnequal   Relation = "unequal"
)

func Sublist(l1, l2 []int) Relation {
	var ok int
	if len(l1) == 0 && len(l2) > 0 {
		return RelationSublist
	} else if len(l1) > 0 && len(l2) == 0 {
		return RelationSuperlist
	} else if len(l1) < len(l2) {
		for i := 0; i < len(l2); i++ {
			for j := 0; j < len(l1); j++ {
				if l1[j] != l2[i+j] {
					break
				} else {
					ok++
				}
			}
			if ok == len(l1) {
				return RelationSublist
			}
			ok = 0
		}
		return RelationUnequal
	} else if len(l1) > len(l2) {
		for i := 0; i < len(l1); i++ {
			for j := 0; j < len(l2); j++ {
				if l2[j] != l1[i+j] {
					break
				} else {
					ok++
				}
			}
			if ok == len(l2) {
				return RelationSuperlist
			}
			ok = 0
		}
		return RelationUnequal
	} else if len(l1) == len(l2) {
		for i := 0; i < len(l1); i++ {
			if l1[i] != l2[i] {
				return RelationUnequal
			}
		}
	}
	return RelationEqual
}

func main() {
	listOne := []int{1, 2, 5}
	listTwo := []int{0, 1, 2, 3, 1, 2, 5, 6}
	fmt.Println(Sublist(listOne, listTwo))
}
