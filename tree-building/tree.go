package tree

import "fmt"

// Record represents the relation between node and parent node.
type Record struct {
	ID     int
	Parent int
}

// Node represents a node of the tree structure.
type Node struct {
	ID       int
	Children []*Node
}

// Build creates a tree structure.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// node ID is a uniq number from 0 to n
	// so it can be represented by an array index
	// the value of this element is the parent ID
	childsID := make([]int, len(records))
	// each record is a Node in the tree structure
	nodes := make([]*Node, len(records))

	for _, r := range records {
		err := checkErrors(r, len(records), nodes)
		if err != nil {
			return nil, err
		}

		childsID[r.ID] = r.Parent
		nodes[r.ID] = &Node{ID: r.ID}
	}

	for id, parent := range childsID {
		if id == 0 {
			continue
		}

		nodes[parent].Children = append(nodes[parent].Children, nodes[id])
	}

	return nodes[0], nil
}

func checkErrors(r Record, size int, nodes []*Node) error {
	if r.ID == 0 && r.Parent != 0 {
		return fmt.Errorf("invalid root node")
	}
	if r.ID != 0 && r.ID == r.Parent {
		return fmt.Errorf("cycle directly")
	}
	if r.ID >= size || r.Parent >= size {
		return fmt.Errorf("invalid ID or ParentID")
	}
	if r.ID < r.Parent {
		return fmt.Errorf("ID must be greater than Parent")
	}
	if nodes[r.ID] != nil {
		return fmt.Errorf("duplicate node")
	}

	return nil
}
