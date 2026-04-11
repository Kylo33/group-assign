package flow

import (
	"math"
)

func (g *Graph) MaxFlow() {
	for {
		seen := make([]bool, g.numVerticies+sourceSinkShift)
		if maxFlowDfs(g, Source, math.MaxInt, seen) < 0 {
			break
		}
	}
}

func maxFlowDfs(g *Graph, node, minFlow int, seen []bool) int {
	if seen[node+sourceSinkShift] {
		return -1
	}
	seen[node+sourceSinkShift] = true

	if node == Sink {
		return minFlow
	}
	
	for _, edge := range g.OutgoingEdges(node) {
		potential := edge.Capacity - edge.Flow
		if potential > 0 {
			augmentBy := maxFlowDfs(g, edge.Dst, min(minFlow, potential), seen)
			if augmentBy < 0 {
				continue
			}

			edge.Flow += augmentBy
			return augmentBy
		}
	}

	for _, edge := range g.IncomingEdges(node) {
		potential := edge.Flow

		if potential > 0 {
			augmentBy := maxFlowDfs(g, edge.Dst, min(minFlow, potential), seen)
			if augmentBy < 0 {
				continue
			}

			edge.Flow -= augmentBy
			return augmentBy
		}
	}

	return -1
}
