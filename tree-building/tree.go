package tree

import (
	"fmt"
	"sort"
)

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

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node, len(records))

	for _, r := range records {
		if err := checkErrors(r, len(records), nodes); err != nil {
			return nil, err
		}

		if r.ID == 0 {
			nodes[0] = &Node{ID: 0}
			continue
		}

		// current node doesn't exist
		if _, ok := nodes[r.ID]; !ok {
			nodes[r.ID] = &Node{ID: r.ID}
		}

		// parent node doesn't exist
		if _, ok := nodes[r.Parent]; !ok {
			nodes[r.Parent] = &Node{ID: r.Parent}
		}

		nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
	}

	return nodes[0], nil
}

func checkErrors(r Record, size int, nodes map[int]*Node) error {
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
	if _, ok := nodes[r.ID]; ok {
		return fmt.Errorf("duplicate node")
	}

	return nil
}
