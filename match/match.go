package match

import (
	"log"

	"github.com/Kylo33/group-assign/flow"
)

type Match[L any, R any] struct {
	From L
	To   []R
}

func Fair[L any, R any](left []L, right []R, coverage int) ([]Match[L, R], error) {
	if coverage > len(right) {
		log.Fatalf("Coverage is larger than items that can be assigned toa")
	}

	graph := flow.NewGraph(len(left) + len(right))

	for leftIndex := range left {
		graph.AddEdge(flow.Source, leftIndex, coverage)
		for rightIndex := range right {
			graph.AddEdge(leftIndex, len(left)+rightIndex, 1)
		}
	}

	// Two passes.
	// 1. allowedRight = ⌊leftCount * coverage / rightCount⌋, run max flow
	// 2. allowedRight = ⌈leftCount * coverage / rightCount⌉, run max flow

	// First pass
	allowedRight := (len(left)*coverage) / len(right)
	for rightIndex := range right {
		graph.AddEdge(len(left)+rightIndex, flow.Sink, allowedRight)
	}
	graph.MaxFlow()

	// Second pass
	allowedRight = (len(left)*coverage + len(right) - 1) / len(right)
	for _, edge := range graph.IncomingEdges(flow.Sink) {
		edge.Capacity = allowedRight
	}
	graph.MaxFlow()
	
	// Find the matches
	var matches []Match[L, R]
	for leftIndex, leftThing := range left {
		match := Match[L, R]{From: leftThing}

		for _, edge := range graph.OutgoingEdges(leftIndex) {
			if edge.Dst == flow.Source {
				continue
			}

			if edge.Flow > 0 {
				match.To = append(match.To, right[edge.Dst - len(left)])
			}
		}

		matches = append(matches, match)
	}

	return matches, nil
}
