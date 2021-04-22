package main

import (
    "fmt"
    "math"
    "gonum.org/v1/gonum/graph/simple"
)

func main() {
    self := 0.0                   // the cost of self connection
    absent := math.Inf(1)       // the wieght returned for absent edges
    
    graph := simple.NewWeightedUndirectedGraph(self, absent)
    fmt.Printf("%v\n", graph)

    var id int64
    //var node simple.Node

    id = 0
    from := simple.Node(id)
    graph.AddNode(from)

    id = 1
    to := simple.Node(id)
    graph.AddNode(to)

    id = 2
    from2 := simple.Node(id)
    graph.AddNode(from2)

    id = 3
    to2 := simple.Node(id)
    graph.AddNode(to2)

    nodeA := graph.Node(int64(2))



    fmt.Printf("%v\n", graph)

    nodes := graph.Nodes()
    fmt.Printf("%v\n", nodes)
    fmt.Printf("%v\n", nodeA)

    weight := float64(40)
    edge := graph.NewWeightedEdge(from, to, weight)
    graph.SetWeightedEdge(edge)

    edge2 := graph.NewWeightedEdge(from2, to2, weight)
    graph.SetWeightedEdge(edge2)

    fmt.Printf("%v\n", graph)
    edges := graph.Edges()
    fmt.Printf("%v\n", edges)

    edge_ := graph.Edge(int64(0) ,int64(1))
    fmt.Printf("%v\n", edge_)

/*
    from := simple.Node(int64(0))
    to := simple.Node(int64(1))
    weight := float64(40)
    graph.AddNode(from)
    graph.AddNode(to)
    
    fmt.Printf("%v\n", graph)
    // Pattern 1 (directly add edge to graph)
   // _ = graph.NewEdge(from, to)
    _ = graph.NewWeightedEdge(from, to, weight)

    fmt.Printf("%v\n", graph)

    edge := graph.Edge(int64(0), int64(1))
    fmt.Printf("%v\n", edge)

// or you can add node without specifying id
//node = graph.NewNode()
//    fmt.Printf("%v\n", node)
*/

}
