package binarysearch

// SearchInts implements the binary search algorithm.
func SearchInts(arr []int, key int) int {
	min, max := 0, len(arr)-1

	middle := (min + max) / 2
	for min <= max {
		if arr[middle] == key {
			return middle
		} else if arr[middle] > key {
			max = middle - 1
		} else {
			min = middle + 1
		}

		middle = (min + max) / 2
	}

	return -1
}
