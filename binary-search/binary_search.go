package binarysearch

// SearchInts implements the binary search algorithm.
func SearchInts(arr []int, key int) int {
	var middle int
	var found bool
	var i int
	for len(arr) > 0 {
		middle = len(arr) / 2
		if arr[middle] == key {
			i += middle
			found = true
			break
		}
		if arr[middle] > key {
			arr = arr[:middle]
			continue
		}
		if arr[middle] < key {
			i += middle + 1
			arr = arr[middle+1:]
			continue
		}
	}

	if !found {
		i = -1
	}

	return i
}
