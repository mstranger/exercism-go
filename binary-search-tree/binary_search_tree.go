package binarysearchtree

// SearchTreeData implements Binary Search Tree structure.
type SearchTreeData struct {
	data  int
	left  *SearchTreeData
	right *SearchTreeData
}

// Bst creates new tree.
func Bst(n int) *SearchTreeData {
	return &SearchTreeData{data: n}
}

// Insert inserts a given number to the binary tree.
func (t *SearchTreeData) Insert(n int) {
	// insert left
	if t.data >= n {
		if t.left == nil {
			t.left = Bst(n)
			return
		} else {
			t.left.Insert(n)
		}
	}
	// insert right
	if t.data < n {
		if t.right == nil {
			t.right = Bst(n)
			return
		} else {
			t.right.Insert(n)
		}
	}
}

// MapStrings returns a sorted array of string type.
// Calls the function f for each element of the tree.
func (t *SearchTreeData) MapString(f func(int) string) []string {
	arr := make([]interface{}, 0)
	traverse(t, &arr)

	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = f(v.(int))
	}

	return res
}

// MapInt return a sorted array of int type.
// Calls the function f for each element of the tree.
func (t *SearchTreeData) MapInt(f func(int) int) []int {
	arr := make([]interface{}, 0)
	traverse(t, &arr)

	res := make([]int, len(arr))
	for i, v := range arr {
		res[i] = f(v.(int))
	}

	return res
}

// LNR tree traversal (in-order)
// returns data in sorted order
func traverse(t *SearchTreeData, arr *[]interface{}) {
	if t == nil {
		return
	}
	traverse(t.left, arr)
	*arr = append(*arr, t.data)
	traverse(t.right, arr)
}
