package graph

type WeightedGraph [][]float64

func FloydWarshall(graph WeightedGraph) WeightedGraph {
	if len(graph) == 0 {
		return nil
	}

	for i := 0; i < len(graph); i++ {
		if len(graph) != len(graph[i]) {
			return nil
		}
	}

	v := len(graph)

	result := make([][]float64, v)

	for i := 0; i < v; i++ {
		result[i] = make([]float64, v)
		for j := 0; j < v; j++ {
			result[i][j] = graph[i][j]
		}
	}

	for k := 0; k < v; k++ {
		for i := 0; i < v; i++ {
			for j := 0; j < v; j++ {
				if result[i][j] > result[i][k]+result[k][j] {
					result[i][j] = result[i][k] + result[k][j]
				}
			}
		}
	}

	return result
}
