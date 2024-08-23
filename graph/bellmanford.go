package graph

import (
	"errors"
	"math"
)

func (g *Graph) BellmanFord(start, end int) (int, bool, error) {
	INF := math.Inf(1)
	distances := make([]float64, g.vertices)

	for i := 0; i < g.vertices; i++ {
		distances[i] = INF
	}
	distances[start] = 0

	for i := 0; i < g.vertices-1; i++ {
		for u, neighbors := range g.edges {
			for v, d := range neighbors {
				thisDist := distances[u] + float64(d)
				if thisDist < distances[v] {
					distances[v] = thisDist
				}
			}
		}
	}

	for u, neighbors := range g.edges {
		for v, d := range neighbors {
			thisDist := distances[u] + float64(d)
			if thisDist < distances[v] {
				return -1, false, errors.New("negative weight cycle exists")
			}
		}
	}

	if distances[end] == INF {
		return -1, false, nil
	}
	return int(distances[end]), true, nil
}
