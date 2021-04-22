package main

import (
    "fmt"
    "os"
    "io"
    "math"
    "strconv"
    "encoding/csv"
    "gonum.org/v1/gonum/graph/simple"
)

func main() {

	// Define Graph
    var self float64 = 0
    var absent float64 = math.Inf(1)
    graph := simple.NewWeightedUndirectedGraph(self, absent)

    // Open File
    fp, err := os.Open("./sample.edgelist")
    if err != nil {
       panic(err)
    }
    defer fp.Close()

    // Read rows and add edges
    reader := csv.NewReader(fp)
    reader.Comma = '\t'
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }

       	id_from, _ := strconv.ParseInt(record[0], 10, 64)
       	id_to, _ := strconv.ParseInt(record[1], 10, 64)
       	weight, _ := strconv.ParseFloat(record[2], 64)
       	
       	weightedEdge := simple.WeightedEdge{F: simple.Node(id_from), T: simple.Node(id_to), W: weight}
        graph.SetWeightedEdge(weightedEdge)
    }

    // Print all the edges
    fmt.Print(graph.WeightedEdges())
}
