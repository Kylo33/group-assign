package flow

const Source = -2
const Sink = -1

// Use this to shift indices up by 2 internally to give space
// for the source and sink vertices.
const sourceSinkShift = 2

type Graph struct {
	numVerticies int
	outgoing [][]*Edge
	incoming [][]*Edge
}

type Edge struct {
	Src      int
	Dst      int
	Capacity int
	Flow     int
}

func NewGraph(numVertices int) Graph {
	return Graph{
		numVerticies: numVertices,
		outgoing: make([][]*Edge, numVertices + sourceSinkShift),
		incoming: make([][]*Edge, numVertices + sourceSinkShift),
	}
}

func (g *Graph) AddEdge(src, dst, capacity int) {
	newEdge := Edge{
		Src: src,
		Dst: dst,
		Capacity: capacity,
	}

	g.outgoing[src + sourceSinkShift] = append(g.outgoing[src + sourceSinkShift], &newEdge)
}

func (g *Graph) IncomingEdges(vertex int) []*Edge {
	return g.incoming[vertex + sourceSinkShift]
}

func (g *Graph) OutgoingEdges(vertex int) []*Edge {
	return g.outgoing[vertex + sourceSinkShift]
}

