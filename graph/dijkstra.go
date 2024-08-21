package graph

import (
	"github.com/homearchbishop/algorithms-golang/structure"
)

type NodeItem struct {
	node int
	dist int
}

func (nc NodeItem) Id() int {
	return nc.node
}

func (nc NodeItem) LargerThan(nc2 interface{}) bool {
	return nc.dist < nc2.(NodeItem).dist // turn MaxHeap to MinHeap
}

func (g *Graph) Dijkstra(start, end int) (int, bool) {
	visited := make(map[int]bool)
	nodes := make(map[int]*NodeItem)
	pq := structure.MaxHeap[NodeItem]{} // actually MinHeap

	pq.Init(nil)

	startNC := &NodeItem{node: start, dist: 0}
	nodes[start] = startNC
	pq.Push(*startNC)

	for pq.Size() > 0 {
		cur := *pq.Pop()
		if cur.node == end {
			break
		}

		visited[cur.node] = true
		for to, d := range g.edges[cur.node] {
			if visited[to] {
				continue
			}
			thisDist := nodes[cur.node].dist + d
			if nodes[to] == nil {
				nodes[to] = &NodeItem{node: to, dist: thisDist}
				pq.Push(*nodes[to])
			} else if nodes[to].dist > thisDist {
				nodes[to].dist = thisDist
				pq.Update(*nodes[to])
			}
		}
	}

	item := nodes[end]
	if item == nil {
		return -1, false
	}
	return item.dist, true
}
