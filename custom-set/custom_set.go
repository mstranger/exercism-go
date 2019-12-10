package stringset

import (
	"fmt"
	"sort"
	"strings"
)

// Set represets a custom set type.
type Set []string

// print set as {"a", "b", "c"}
func (s Set) String() string {
	t := []string{}
	for _, v := range s {
		vf := fmt.Sprintf("\"%s\"", v)
		t = append(t, vf)
	}

	return fmt.Sprintf("{%s}", strings.Join(t, ", "))
}

// New returns a new empty set.
func New() Set {
	return Set{}
}

// NewFromSlice returns new uniq set from the given slice.
func NewFromSlice(slice []string) Set {
	return Set(uniq(slice))
}

// Equal check if the given sets are equal (with equal elements).
func Equal(s1, s2 Set) bool {
	t1, t2 := sortInOrder(s1, s2)
	return strings.Join(t1, "") == strings.Join(t2, "")
}

// Subset checks if one set is a subset of the second.
func Subset(s1, s2 Set) bool {
	t1, t2 := sortInOrder(s1, s2)
	j1, j2 := strings.Join(t1, ""), strings.Join(t2, "")

	if j2 == "" && j1 != "" {
		return false
	}

	return strings.Contains(j1, j2) || strings.Contains(j2, j1)
}

// Disjoint check if the given sets are disjoint sets.
func Disjoint(s1, s2 Set) bool {
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v2 == v1 {
				return false
			}
		}
	}

	return true
}

// Intersection finds intersection of sets.
func Intersection(s1, s2 Set) Set {
	inter := []string{}
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v2 == v1 {
				inter = append(inter, v2)
				continue
			}
		}
	}

	return NewFromSlice(inter)
}

// Difference finds the difference of two sets.
func Difference(s1, s2 Set) Set {
	diff := []string{}

	for _, v1 := range s1 {
		var mark bool
		for _, v2 := range s2 {
			if v2 == v1 {
				mark = true
				break
			}
		}

		if !mark {
			diff = append(diff, v1)
		}
	}

	return NewFromSlice(diff)
}

// Union creates new union from the given sets.
func Union(s1, s2 Set) Set {
	joined := append(s1, s2...)
	return NewFromSlice(joined)
}

// Add adds a new element to the set (if not exists).
func (s *Set) Add(e string) {
	if s.Has(e) {
		return
	}
	*s = append(*s, e)
}

// IsEmpty checks is the given set empty.
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has checks uniqueness of the given element in the set.
func (s Set) Has(e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

// helper func, delete duplicate items from the given slice
func uniq(sl []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, v := range sl {
		if _, ok := keys[v]; !ok {
			keys[v] = true
			list = append(list, v)
		}
	}
	return list
}

// helper func, return two sorted strings bases on the given sets
func sortInOrder(s1, s2 Set) ([]string, []string) {
	t1 := make([]string, len(s1))
	t2 := make([]string, len(s2))
	copy(t1, s1)
	copy(t2, s2)
	sort.Strings(t1)
	sort.Strings(t2)
	return t1, t2
}
