package sublist

// Relation is a ralation between two lists.
type Relation string

// Sublist determines if the first list is contained within the second list.
func Sublist(l1, l2 []int) Relation {
	switch {
	case len(l1) < len(l2) && isSub(l1, l2):
		return "sublist"
	case len(l1) > len(l2) && isSub(l2, l1):
		return "superlist"
	case equal(l1, l2):
		return "equal"
	default:
		return "unequal"
	}
}

// check if the first list is a sublist for the second list
func isSub(l1, l2 []int) bool {
	for i := 0; i <= len(l2)-len(l1); i++ {
		if equal(l1, l2[i:i+len(l1)]) {
			return true
		}
	}
	return false
}

// compare two lists
func equal(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i, v := range l2 {
		if v != l1[i] {
			return false
		}
	}
	return true
}
