package graph

type Graph struct {
	edges    map[int]map[int]int
	Directed bool
}

func (g *Graph) AddVertex(v int) {
	if g.edges == nil {
		g.edges = make(map[int]map[int]int)
	}

	if _, ok := g.edges[v]; !ok {
		g.edges[v] = make(map[int]int)
	}
}

func (g *Graph) AddEdge(one, two int) {
	g.AddWeightedEdge(one, two, 0)
}

func (g *Graph) AddWeightedEdge(one, two, weight int) {
	g.AddVertex(one)
	g.AddVertex(two)

	g.edges[one][two] = weight
	if !g.Directed {
		g.edges[two][one] = weight
	}
}
