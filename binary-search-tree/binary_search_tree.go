// package main
package binarysearchtree

import (
	"sort"
)

type SearchTreeData struct {
	data  int
	left  *SearchTreeData
	right *SearchTreeData
}

func Bst(n int) *SearchTreeData {
	return &SearchTreeData{data: n}
}

func (t *SearchTreeData) Insert(n int) {
	if t == nil {
		return
	}
	if t.data >= n {
		if t.left == nil {
			t.left = &SearchTreeData{data: n}
			return
		} else {
			t.left.Insert(n)
		}
	}
	if t.data < n {
		if t.right == nil {
			t.right = &SearchTreeData{data: n}
			return
		} else {
			t.right.Insert(n)
		}
	}
}

func (t *SearchTreeData) MapString(f func(int) string) []string {
	arr := make([]interface{}, 0)
	traverse(t, &arr)

	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = f(v.(int))
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}

func (t *SearchTreeData) MapInt(f func(int) int) []int {
	arr := make([]interface{}, 0)
	traverse(t, &arr)

	res := make([]int, len(arr))
	for i, v := range arr {
		res[i] = f(v.(int))
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}

func traverse(t *SearchTreeData, arr *[]interface{}) {
	if t == nil {
		return
	}
	*arr = append(*arr, t.data)
	traverse(t.left, arr)
	traverse(t.right, arr)
	return
}

// func main() {
// 	t := Bst(4)
// 	t.Insert(2)
// 	t.Insert(6)
// 	t.Insert(3)
// 	t.Insert(1)
// 	t.Insert(5)
// 	t.Insert(7)

// 	arr1 := t.MapInt(func(x int) int { return x })
// 	arr2 := t.MapString(strconv.Itoa)

// 	fmt.Printf("%v\n", arr1)
// 	fmt.Printf("%q\n", arr2)
// }
