package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type ByParent []Record
type ByID []Record

func (r ByParent) Len() int      { return len(r) }
func (r ByParent) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r ByParent) Less(i, j int) bool {
	if r[i].Parent == r[j].Parent {
		return r[i].ID < r[j].ID
	}

	return r[i].Parent < r[j].Parent
}

func (r ByID) Len() int           { return len(r) }
func (r ByID) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByID) Less(i, j int) bool { return r[i].ID < r[j].ID }

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	n := &Node{}

	sort.Sort(ByID(records))
	if t := checkErrors(records); t == true {
		return nil, fmt.Errorf("error")
	}

	// sort records
	sort.Sort(ByParent(records))
	if records[0].ID != 0 {
		return nil, fmt.Errorf("no root node")
	}

	// if t := checkErrors(sort.Sort(ByID(records))); t == true {
	// 	return nil, fmt.Errorf("error")
	// }

	// fmt.Println(records)

	for _, v := range records[1:] {
		if v.Parent >= v.ID {
			return nil, fmt.Errorf("invalid ID and ParentID numbers")
		}

		if t := findNodeWithId(n, v.ID); t != nil {
			return nil, fmt.Errorf("duplicate ID")
		}

		parent := findNodeWithId(n, v.Parent)
		if parent != nil {
			parent.Children = append(parent.Children, &Node{ID: v.ID})
		}

		// fmt.Println(parent)
	}

	// fmt.Println(n)
	return n, nil
}

func checkErrors(sorted []Record) bool {
	// fmt.Println(sorted)
	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i].ID+1 != sorted[i+1].ID {
			return true
		}
	}

	return false
}

func findNodeWithId(node *Node, id int) *Node {
	// fmt.Println(id, node)

	if id == 0 {
		return node
	}

	if id == node.ID {
		return node
	}

	for _, v := range node.Children {
		current := findNodeWithId(v, id)
		if current != nil && current.ID == id {
			return current
		}
	}

	// not found
	return nil
}
