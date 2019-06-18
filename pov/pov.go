package pov

// Graph is directed graph of string nodes.
type Graph map[string][]string

// New creates a new empty graph.
func New() *Graph {
	return &Graph{}
}

// AddNode adds a node to the graph.
func (g *Graph) AddNode(nodeLabel string) {
	(*g)[nodeLabel] = make([]string, 0)
}

// AddArc adds an edge to the graph.
func (g *Graph) AddArc(from, to string) {
	children := (*g)[from]
	(*g)[from] = append(children, to)
}

// ArcList displays all edges in the graph.
func (g *Graph) ArcList() []string {
	edges := []string{}
	for from, v := range *g {
		for _, to := range v {
			edges = append(edges, from+" -> "+to)
		}
	}
	return edges
}

// ChangeRoot reroots the graph.
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	path := g.getPath(oldRoot, newRoot)
	for i := 0; i < len(path)-1; i++ {
		oldTo, oldFrom := path[i], path[i+1]
		g.removeArc(oldFrom, oldTo)
		g.AddArc(oldTo, oldFrom)
	}
	return g
}

// remove an egde from the graph
func (g *Graph) removeArc(from, to string) {
	children := (*g)[from]
	newChildren := []string{}
	for _, child := range children {
		if child != to {
			newChildren = append(newChildren, child)
		}
	}
	(*g)[from] = newChildren
}

// find the list from one node to another
func (g *Graph) getPath(from, to string) []string {
	if from == to {
		return []string{to}
	}
	for _, child := range (*g)[from] {
		if path := g.getPath(child, to); path != nil {
			return append(path, from)
		}
	}
	return nil
}
