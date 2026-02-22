package sublist

// Relation type is defined in relations.go file.

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
		ok = 0
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
		ok = 0
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
