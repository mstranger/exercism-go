package listops

// IntList is a slice of integers
type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Length returns the length of the given list.
func (list IntList) Length() (len int) {
	for range list {
		len++
	}
	return
}

// Reverse returns the new list with the reversed values.
func (list IntList) Reverse() IntList {
	newList := make(IntList, list.Length())
	for i := range list {
		newList[i] = list[list.Length()-i-1]
	}
	return newList
}

// Append appends the given list to the calling list.
func (list IntList) Append(src IntList) IntList {
	len := list.Length() + src.Length()
	newList := make(IntList, len)
	curr := 0
	for ; curr < list.Length(); curr++ {
		newList[curr] = list[curr]
	}
	for i := 0; curr < len; i, curr = i+1, curr+1 {
		newList[curr] = src[i]
	}

	return newList
}

// Concat concatenates all given lists to the calling list.
func (list IntList) Concat(args []IntList) IntList {
	for _, l := range args {
		list = list.Append(l)
	}

	return list
}

// Foldl reduces a list from left to right with initial value.
func (list IntList) Foldl(fn binFunc, init int) int {
	result := init
	for _, v := range list {
		result = fn(result, v)
	}
	return result
}

// Foldr reduces a list from right to left with initial value.
func (list IntList) Foldr(fn binFunc, init int) int {
	result := init
	for i := list.Length() - 1; i >= 0; i-- {
		result = fn(list[i], result)
	}
	return result
}

// Filter returns new list with filtered values by given function.
func (list IntList) Filter(fn predFunc) IntList {
	filtered := make(IntList, 0)

	for _, v := range list {
		if fn(v) {
			filtered = filtered.Append(IntList{v})
		}
	}

	return filtered
}

// Map calls the given function for every entry in the list.
func (list IntList) Map(fn unaryFunc) IntList {
	for i, v := range list {
		list[i] = fn(v)
	}

	return list
}
