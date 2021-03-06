package main

import (
	"fmt"
	yaml "github.com/goccy/go-yaml"
	"io/ioutil"
	"local.packages/graph"
	"net"
)

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

type Data struct {
	Routers []Router `yaml:"routers"`
	Subnets []Subnet `yaml:"subnets"`
}

func main() {

	buf, err := ioutil.ReadFile("./network.yaml")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("buf: %+v\n", string(buf))

	var d Data
	router_dic := make(map[string]Router)
	subnetaddr_dic := make(map[string]Subnet)
	subnetname_dic := make(map[string]Subnet)

	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		panic(err)
	}

	var id int64
	id = 1
	nodes := []graph.NetNode{}
	g := graph.NewGraph(nodes)
	for _, v := range d.Routers {
		//fmt.Printf("%T\n", v)
		//fmt.Printf("%s\n", v.GetName())
		v._ID = id
		id++
		g.AddNode(v)
		router_dic[v.HostName] = v

	}

	for _, v := range d.Subnets {
		//fmt.Printf("%T\n", v)
		//fmt.Printf("%s\n", v.GetName())
		v._ID = id
		id++
		g.AddNode(v)
		subnetaddr_dic[v.Netaddr] = v
		subnetname_dic[v.Name] = v

	}
	//fmt.Printf("----------------------------------------\n")
	//fmt.Printf("router_dic\n%v\n\n", router_dic)
	//fmt.Printf("subnet_dic\n%v\n\n", subnet_dic)
	//fmt.Printf("----------------------------------------\n")

	for _, v := range router_dic {
		for _, i := range v.Interfaces {
			_, ipnet, _ := net.ParseCIDR(i.Ipaddr)
			masklen, _ := ipnet.Mask.Size()
			//fmt.Println(ip) //
			//fmt.Println(ipnet.IP) // 10.0.0.0
			//fmt.Println(ipnet.Mask) // ffffff00
			//fmt.Println(masklen) // 24
			subnet := fmt.Sprintf("%s/%d", ipnet.IP, masklen)
			s := subnetaddr_dic[subnet]
			g.SetEdge(v, s)
		}
	}

	g.CalculateShortest()
        //------------------------------------------------------------------------

	//g.Dump()

        // GetNodes
	iter_nodes := g.GetNodes()

	for iter_nodes.Next() {
		node := iter_nodes.Node()
		switch v := node.(type) {
		case Router:
			fmt.Printf("%s\n", v.HostName)
			fmt.Printf("%s\n", v.GetName())
		case Subnet:
			fmt.Printf("%s\n", v.Name)
			fmt.Printf("%s\n", v.GetName())
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}

        // Node
	var cnt int64 = 1
	for {
		node := g.Node(cnt)
		cnt++
		if node != nil {
			switch v := node.(type) {
			case Router:
				fmt.Printf("%s\n", v.HostName)
				fmt.Printf("%s\n", v.GetName())
			case Subnet:
				fmt.Printf("%s\n", v.Name)
				fmt.Printf("%s\n", v.GetName())
			default:
				fmt.Printf("I don't know about type %T!\n", v)
			}
		} else {
			break
		}

	} 

        // GetEdges
	iter_edges := g.GetEdges()

	for iter_edges.Next() {
		edge := iter_edges.Edge()
		printEdge(edge.From().(graph.NetNode), edge.To().(graph.NetNode))
	}

        // GetNeighbour
	iter_nodes = g.GetNeighbour(router_dic["RA"])
	for iter_nodes.Next() {
		node := iter_nodes.Node()
		printNode(node.(graph.NetNode))
	}

	iter_nodes = g.GetNeighbour(subnetname_dic["subAC"])
	for iter_nodes.Next() {
		node := iter_nodes.Node()
		//fmt.Printf("%v\n", edge.From())
		//fmt.Printf("%v\n", edge.To())
		printNode(node.(graph.NetNode))
	}


        // GetBetween
	//paths := g.GetBitween(router_dic["RA"], router_dic["RY"])
	paths := g.GetBitween(subnetname_dic["subA"], subnetname_dic["subD"])
	fmt.Printf("paths count:%d\n", len(paths))
	for i, array := range paths {
		for j, value := range array {
			fmt.Println(fmt.Sprintf("   [%d][%d] :%v",
				i, j, value.(graph.NetNode).GetName()))
		}
		fmt.Printf("\n")
	}

        // ConnectedComponents
        cons := g.ConnectedComponents()
	fmt.Printf("cons count:%d\n", len(cons))
	for i, array := range cons {
		for j, value := range array {
			fmt.Println(fmt.Sprintf("   [%d][%d] :%v\t\t%d",
				i, j, value.(graph.NetNode).GetName(),
				      value.(graph.NetNode).ID()))
		}
		fmt.Printf("\n")
	}

        // IsPathIn
        path := [] graph.NetNode{}
        path = append( path, router_dic["RA"])
        path = append( path, subnetname_dic["subA"])
        fmt.Printf("%v\n",g.IsPathIn(path))

        path = [] graph.NetNode{}
        path = append( path, subnetname_dic["subA"])
        path = append( path, router_dic["RA"])
        path = append( path, subnetname_dic["subAB"])
        path = append( path, router_dic["RB"])
        fmt.Printf("%v\n",g.IsPathIn(path))

        path = [] graph.NetNode{}
        path = append( path, subnetname_dic["subA"])
        path = append( path, router_dic["RA"])
        path = append( path, subnetname_dic["subAB"])
        path = append( path, router_dic["RC"])
        fmt.Printf("%v\n",g.IsPathIn(path))

        path = [] graph.NetNode{}
        path = append( path, subnetname_dic["subA"])
        path = append( path, router_dic["RA"])
        path = append( path, subnetname_dic["subAB"])
        path = append( path, router_dic["RB"])
        path = append( path, subnetname_dic["subBD"])
        path = append( path, router_dic["RY"])
        path = append( path, subnetname_dic["subY"])
        fmt.Printf("%v\n",g.IsPathIn(path))

        //g.RemoveEdge(router_dic["RY"], subnetname_dic["subY"])
        //g.RemoveNode(subnetname_dic["subY"])
        fmt.Printf("%v\n",g.IsPathIn(path))

}

func printNode(node graph.NetNode) {
	switch v := node.(type) {
	case Router:
		fmt.Printf("Router [%s]", v.GetName())
	case Subnet:
		fmt.Printf("Subnet %s", v.GetName())
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
	fmt.Printf("\n")
}

func printEdge(from graph.NetNode, to graph.NetNode) {
	switch v := from.(type) {
	case Router:
		fmt.Printf("[%s]", v.GetName())
	case Subnet:
		fmt.Printf("%s", v.GetName())
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
	fmt.Printf(" -> ")
	switch v := to.(type) {
	case Router:
		fmt.Printf("[%s]", v.GetName())
	case Subnet:
		fmt.Printf("%s", v.GetName())
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
	fmt.Printf("\n")
}
