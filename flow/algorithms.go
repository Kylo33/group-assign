package flow

import (
	"math"
	"math/rand"
)

func (g *Graph) MaxFlow() {
	for {
		seen := make([]bool, g.numVerticies+sourceSinkShift)
		if maxFlowDfs(g, Source, math.MaxInt, seen) < 0 {
			break
		}
	}
}

type flowCandidate struct {
	incoming  bool
	potential int
	edge      *Edge
}

func maxFlowDfs(g *Graph, node, minFlow int, seen []bool) int {
	if seen[node+sourceSinkShift] {
		return -1
	}
	seen[node+sourceSinkShift] = true

	if node == Sink {
		return minFlow
	}

	candidates := findFlowCandidates(g, node)

	rand.Shuffle(
		len(candidates),
		func(i, j int) { candidates[i], candidates[j] = candidates[j], candidates[i] },
	)

	for _, candidate := range candidates {
		if candidate.potential <= 0 {
			continue
		}

		augmentBy := maxFlowDfs(g, candidate.edge.Dst, min(minFlow, candidate.potential), seen)
		if augmentBy < 0 {
			continue
		}

		if candidate.incoming {
			candidate.edge.Flow -= augmentBy
		} else {
			candidate.edge.Flow += augmentBy
		}
		return augmentBy
	}

	return -1
}

func findFlowCandidates(g *Graph, node int) []flowCandidate {
	var candidates []flowCandidate

	for _, edge := range g.OutgoingEdges(node) {
		potential := edge.Capacity - edge.Flow
		candidates = append(candidates, flowCandidate{incoming: false, potential: potential, edge: edge})
	}

	for _, edge := range g.IncomingEdges(node) {
		potential := edge.Flow
		candidates = append(candidates, flowCandidate{incoming: true, potential: potential, edge: edge})
	}
	return candidates
}

