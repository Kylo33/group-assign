package flow

const Source = -2
const Sink = -1

// Use this to shift indices up by 2 internally to give space
// for the source and sink vertices.
const sourceSinkShift = 2

type Graph struct {
	adj [][]Edge
}

type Edge struct {
	Src      int
	Dst      int
	Capacity int
	Flow     int
}

func NewGraph(numVertices int) Graph {
	return Graph{
		adj: make([][]Edge, numVertices + sourceSinkShift),
	}
}

func (g *Graph) AddEdge(src, dst, capacity int) {
	newEdge := Edge{
		Src: src,
		Dst: dst,
		Capacity: capacity,
	}

	g.adj[src + sourceSinkShift] = append(g.adj[src + sourceSinkShift], newEdge)
}

func (g *Graph) IncidentEdges(vertex int) []Edge {
	return g.adj[vertex + sourceSinkShift]
}

