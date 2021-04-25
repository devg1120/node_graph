package graph

import (
	"encoding/csv"
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/iterator"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
	"math"
	"os"
)

type Graph struct {
	//graph *simple.DirectedGraph
	graph       *simple.WeightedUndirectedGraph
	allShortest *path.AllShortest
}

func NewGraph(nodes []NetNode) Graph {
	self := 0.0           // the cost of self connection
	absent := math.Inf(1) // the wieght returned for absent edges

	graph := Graph{
		//graph: simple.NewDirectedGraph(),
		graph:       simple.NewWeightedUndirectedGraph(self, absent),
		allShortest: nil,
	}

	graph.BuildGraph(nodes)
	return graph
}

func (g *Graph) AddNode(node NetNode) {
	g.graph.AddNode(node)
}

func (g *Graph) SetEdge(from graph.Node, to graph.Node) {
	edge := g.graph.NewWeightedEdge(from, to, float64(1))
	g.graph.SetWeightedEdge(edge)
}

func (g *Graph) RemoveNode(node graph.Node ) {
	g.graph.RemoveNode(node.ID() )
}

func (g *Graph) RemoveEdge(from graph.Node, to graph.Node) {
	g.graph.RemoveEdge(from.ID(), to.ID())
	g.graph.RemoveEdge(to.ID(), from.ID())
}

func (g *Graph) CalculateShortest() {
	allst := path.DijkstraAllPaths(g.graph)
	g.allShortest = &allst
}

func (g *Graph) Node(id int64) graph.Node {
	return g.graph.Node(id)
}

func (g *Graph) GetNodes() graph.Nodes {
	return g.graph.Nodes()
}

func (g *Graph) GetEdges() graph.Edges {
	return g.graph.Edges()
}

//func (g *Graph) GetBetween(node1, node2 graph.Node ) graph.Nodes{
func (g *Graph) GetBitween(node1, node2 graph.Node) [][]graph.Node {
	id1 := node1.ID()
	id2 := node2.ID()
	paths, _ := g.allShortest.AllBetween(id1, id2)
	return paths
}

func (g *Graph) GetNeighbour(node graph.Node) graph.Nodes {
	var neighbour []graph.Node
	edges := g.graph.Edges()
	id := node.ID()
	for edges.Next() {
		edge := edges.Edge()
		if edge.From().ID() == id {
			//fmt.Printf("From OK id:%d\n", id)
			neighbour = append(neighbour, edge.To())
		}
		if edge.To().ID() == id {
			//fmt.Printf("To   OK id:%d\n", id)
			neighbour = append(neighbour, edge.From())
		}
	}
	return iterator.NewOrderedNodes(neighbour)
}

func (g *Graph) BuildGraph(nodes []NetNode) {
	if nodes == nil {
		return
	}
	for _, node := range nodes {
		g.AddNode(node)
	}
}

func (g *Graph) ConnectedComponents() [][]graph.Node {
        paths := topo.ConnectedComponents(g.graph)
	return paths
}

//func (g *Graph) IsPathIn(path []graph.Node) bool {
func (g *Graph) IsPathIn(path []NetNode) bool {
        path_ := []graph.Node{}
        for _, v := range path {
              path_ = append( path_, v.(graph.Node))
        }
        return topo.IsPathIn(g.graph, path_)
	
}
func (g *Graph) PathExistsIn(node1 , node2 graph.Node) bool {
        return topo.PathExistsIn(g.graph, node1, node2)
	
}

func (g *Graph) Dump() {
	nodes := g.graph.Nodes()
	for _, node := range graph.NodesOf(nodes) {
		fmt.Printf("node:%v\n", node)
	}
	edges1 := g.graph.Edges()
	for _, edge := range graph.EdgesOf(edges1) {
		fmt.Printf("edge:%v\n", edge)
	}
	fmt.Printf("%v\n", g.graph.WeightedEdges())
	edges_weight := g.graph.WeightedEdges()
	fp2, err := os.Create("./weighted.edgelist")
	if err != nil {
		panic(err)
	}
	defer fp2.Close()

	writer := csv.NewWriter(fp2)
	writer.Comma = '\t'

	for edges_weight.Next() {
		var record []string
		r := edges_weight.WeightedEdge().From()
		s := edges_weight.WeightedEdge().To()
		fmt.Printf("r %T\n", r)
		fmt.Printf("r %v\n", r)
		fmt.Printf("s %T\n", s)
		//fmt.Printf("r %s\n", r.GetName())
		record = append(record, fmt.Sprint(edges_weight.WeightedEdge().From()))
		record = append(record, fmt.Sprint(edges_weight.WeightedEdge().To()))
		//record = append(record, fmt.Sprint(edges_weight.WeightedEdge().Weight()))
		writer.Write(record)
	}
	writer.Flush()
}
