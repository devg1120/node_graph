package main

import "local.packages/graph"

/*
                  subA                     subB
               +-------+                +-------+
                   |                        |
                   |         subAB          |
                 [RA]----------------------[RB]
subX               |                        |                subY
+                  |                        |                 +
|                  |                        |                 |
|------[RX]------- |subAC              subBD|--------[RY]-----|
|                  |                        |                 |
+                  |                        |                 +
                   |                        |
                 [RC]----------------------[RD]
                   |         subCD          |
                   |                        |
               +-------+                +-------+
                  subC                     subD
     
*/

func main() {

 nodes := []graph.Node{}
 router :=  Router{ "routerA",1 }
 router2 :=  Router{ "routerB",2 }

 subnet :=  Subnet{ "subnet1", 10}

 nodes = append(nodes, router)
// nodes = append(nodes, subnet)
 g := graph.NewGraph(nodes)
 g.AddNode(subnet)
 g.AddNode(router2)


 g.SetEdge(router, subnet)
 g.SetEdge(router2, subnet)

 g.Dump()

// weight := float64(40)
// edge := g.NewWeightedEdge(router, subnet, weight)
// g.SetWeightedEdge(edge)

// fmt.Printf("%v\n", g)
// g.Dump()

/*
    self := 0.0                   // the cost of self connection
    absent := 10.0       // the wieght returned for absent edges

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
*/
}
