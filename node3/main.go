package main

import (
	"fmt"
	"local.packages/graph"
	"net"
	//"strings"
	"io/ioutil"
	//yaml "gopkg.in/yaml.v2"
	yaml "github.com/goccy/go-yaml"
	//  "gopkg.in/go-playground/validator.v9"
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
	subnet_dic := make(map[string]Subnet)

	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		panic(err)
	}

	var id int64
	id = 1
	nodes := []graph.NetNode{}
	g := graph.NewGraph(nodes)
	for _, v := range d.Routers {

		fmt.Printf("%T\n", v)
		fmt.Printf("%s\n", v.GetName())
		v._ID = id
		id++
		//nodes = append(nodes, v)
		g.AddNode(v)
		router_dic[v.HostName] = v

	}

	for _, v := range d.Subnets {

		fmt.Printf("%T\n", v)
		fmt.Printf("%s\n", v.GetName())
		v._ID = id
		id++
		//nodes = append(nodes, v)
		g.AddNode(v)
		subnet_dic[v.Netaddr] = v

	}
	//fmt.Printf("----------------------------------------\n")
	//fmt.Printf("router_dic\n%v\n\n", router_dic)
	//fmt.Printf("subnet_dic\n%v\n\n", subnet_dic)
	//fmt.Printf("----------------------------------------\n")

	//fmt.Printf("----------------------------------------\n")
	for _, v := range router_dic {

		//fmt.Printf("%s\n", v.HostName)
		for _, i := range v.Interfaces {
			//fmt.Printf("    %s\t%s\n", i.Name, i.Ipaddr)
			//ip, ipnet, _ := net.ParseCIDR(i.Ipaddr)
			_, ipnet, _ := net.ParseCIDR(i.Ipaddr)
			masklen, _ := ipnet.Mask.Size()
			//fmt.Println(ip) //
			//fmt.Println(ipnet.IP) // 10.0.0.0
			//fmt.Println(ipnet.Mask) // ffffff00
			//fmt.Println(masklen) // 24
			subnet := fmt.Sprintf("%s/%d", ipnet.IP, masklen)
			//fmt.Println(subnet) //
			s := subnet_dic[subnet]
			//fmt.Printf("r   %v\n", v)
			//fmt.Printf("s   %v\n", s)
			g.SetEdge(v, s)

		}

	}

	g.Dump()

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

	//node := g.Node(int64(2))

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

	} //end for

	iter_edges := g.GetEdges()

	for iter_edges.Next() {
		edge := iter_edges.Edge()
		//fmt.Printf("%v\n", edge.From())
		//fmt.Printf("%v\n", edge.To())
                printEdge(edge.From().(graph.NetNode), edge.To().(graph.NetNode))
	}

	//connect_nodes := g.From(int64(5))
        //fmt.Printf("%v\n", connect_nodes)
        iter_nodes = g.GetNeighbour( router_dic["RA"])
	for iter_nodes.Next() {
		node := iter_nodes.Node()
		//fmt.Printf("%v\n", edge.From())
		//fmt.Printf("%v\n", edge.To())
                printNode(node.(graph.NetNode))
	}
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
