package main

import (
     "fmt"
    "gonum.org/v1/gonum/graph"
    "gonum.org/v1/gonum/graph/multi"
)

type Int64s map[int64]struct{}

func (set Int64s) Add(e int64) {
    set[e] = struct{}{}
}

func (set Int64s) Has(e int64) bool {
    _, ok := set[e]
    return ok
}

type myLine struct {
    line graph.Line
    direction bool
}

func Degree(g *multi.UndirectedGraph, nid int64) (deg int) {
    m := g.From(nid)
    fac := map[bool]int{true: 1, false: 2}  // {ordinary, self-loop}
    for m.Next() {
        mid := m.Node().ID()
        deg += g.Lines(nid, mid).Len() * fac[nid != mid]
    }
    return deg
}

func connectingPaths(g *multi.UndirectedGraph) (nodes Int64s, paths [][]myLine) {
// nodes with their degree not being 2
    nodes = make(Int64s)
    n := g.Nodes()
    for n.Next() {
        nid := n.Node().ID()
        if Degree(g, nid) != 2 {nodes.Add(nid)}
    }

// paths among the nodes
    linesVisited := make(Int64s)
    for uid, _ := range nodes {
        v := g.From(uid)
        for v.Next() {
            vid := v.Node().ID()
            luv := g.Lines(uid, vid)
            for luv.Next() {
                line := luv.Line()
                if linesVisited.Has(line.ID()) {continue}
                linesVisited.Add(line.ID())
                path := []myLine{myLine{line, uid==line.From().ID()}}
                xid := vid
                NEXT:
                for {
                    if nodes.Has(xid) {break}
                    w := g.From(xid)
                    for w.Next() {
                        lxw := g.Lines(xid, w.Node().ID())
                        for lxw.Next() {
                            line = lxw.Line()
                            if !linesVisited.Has(line.ID()) {
                                linesVisited.Add(line.ID())
                                path = append(path, myLine{line, xid==line.From().ID()})
                                xid = w.Node().ID()
                                goto NEXT
                            }
                        }
                    }
                }
                paths = append(paths, path)
            }
        }
    }
    return nodes, paths
}

func printPaths(paths [][]myLine) {
    for j, path := range paths {
        fmt.Printf("%d: nodes: (", j)
        for i, p := range path {
            f, t := p.line.From().ID(), p.line.To().ID()
            if !p.direction {f, t = t, f}
            if i == 0 {fmt.Printf("%d", f)}
            fmt.Printf(" %d", t)
        }
        fmt.Printf(")  lines: {")
        for i, p := range path {
            if i != 0 {fmt.Printf(" ")}
            fmt.Printf("%d", p.line.ID())
        }
        fmt.Printf("}\n")
    }
}

func generateGraph() *multi.UndirectedGraph {
	lines := []struct{ F, T int64 }{
        {0, 1}, {1, 2}, {1, 2}, {1, 3}, {3, 4}, {4, 2}, {2, 5}, {5, 6}, {7, 8}, {8, 9}, {9, 10}, {9, 9},
    }

    g := multi.NewUndirectedGraph()
    for i, l := range lines {
        g.SetLine(multi.Line{F: multi.Node(l.F), T: multi.Node(l.T), UID: int64(i)})
    }

    return g
}

func main() {
    g := generateGraph()
    nodes, paths := connectingPaths(g)
    fmt.Printf("len(nodes, paths) = (%d, %d)\n", len(nodes), len(paths))
    printPaths(paths)
    return
}
