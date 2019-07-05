package sublist

// Relation is a ralation between two lists.
type Relation string

// Sublist determines if the first list is contained within the second list.
func Sublist(l1, l2 []int) Relation {
	order := true
	// the first list should be shorter
	if len(l2) < len(l1) {
		l1, l2 = l2, l1
		order = false
	}

	if len(l1) == len(l2) && equal(l1, l2) {
		return "equal"
	}

	// empty list
	if len(l1) == 0 {
		if order {
			return "sublist"
		}
		return "superlist"
	}

	ids := findAllIdx(l2, l1[0])

	if len(ids) == 0 {
		return "unequal"
	}

	for _, i := range ids {

		if i+len(l1) > len(l2) {
			return "unequal"
		}

		sub := l2[i : i+len(l1)]

		b := equal(l1, sub)

		if b && order {
			return "sublist"
		}

		if b && !order {
			return "superlist"
		}

	}

	return "unequal"
}

// find all positions for the given value
func findAllIdx(l []int, e int) []int {
	ids := make([]int, 0)

	for i, v := range l {
		if v == e {
			ids = append(ids, i)
		}
	}

	return ids
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
