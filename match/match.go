package match

import (
	"log"
	"math/rand"

	"github.com/Kylo33/group-assign/flow"
)

type Match[L any, R any] struct {
	From L
	To   []R
}

func Fair[L any, R any](left []L, right []R, coverage int) []Match[L, R] {
	if coverage > len(right) {
		log.Fatalf("Coverage is larger than items that can be assigned toa")
	}

	graph := flow.NewGraph(len(left) + len(right))

	for leftIndex := range left {
		// Add indices from source → left nodes with weight (coverage)
		graph.AddEdge(flow.Source, leftIndex, coverage)
		// Add indices from left nodes → right nodes with weight 1
		for rightIndex := range right {
			graph.AddEdge(leftIndex, len(left)+rightIndex, 1)
		}
	}

	// Decide the number of problems per person
	assignedRightBase := (len(left) * coverage) / len(right)
	assignedRight := make([]int, len(right))
	for i := range right {
		assignedRight[i] = assignedRightBase
		if i < (len(left)*coverage)%len(right) {
			assignedRight[i]++
		}
	}
	rand.Shuffle(len(assignedRight), func(i, j int) { assignedRight[i], assignedRight[j] = assignedRight[j], assignedRight[i] })

	for rightIndex := range right {
		graph.AddEdge(len(left)+rightIndex, flow.Sink, assignedRight[rightIndex])
	}
	graph.RandomMaxFlow()

	// Find the matches
	var matches []Match[L, R]
	for leftIndex, leftThing := range left {
		match := Match[L, R]{From: leftThing}

		for _, edge := range graph.OutgoingEdges(leftIndex) {
			if edge.Dst == flow.Source {
				continue
			}

			if edge.Flow > 0 {
				match.To = append(match.To, right[edge.Dst-len(left)])
			}
		}

		matches = append(matches, match)
	}

	return matches
}
