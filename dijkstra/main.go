package main

import (
    "fmt"
    "os"
    "io"
    "math"
    "strconv"
    "encoding/csv"
    "gonum.org/v1/gonum/graph/simple"
    "gonum.org/v1/gonum/graph/path"
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

    // calculate shortest path
    allShortest := path.DijkstraAllPaths(graph)

    fmt.Print("from 1 to 18: ")
    fmt.Print(allShortest.AllBetween(int64(1), int64(18)))
    fmt.Print("\n")
    fmt.Print("from 3 to 15: ")
    fmt.Print(allShortest.AllBetween(int64(3), int64(15)))
    fmt.Print("\n")
    fmt.Print("from 6 to 17: ")
    fmt.Print(allShortest.AllBetween(int64(6), int64(17)))
}
