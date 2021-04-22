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
	nodes := []graph.Node{}
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
                fmt.Printf("%T\n", iter_nodes)
                fmt.Printf("%v\n", iter_nodes)
                r := Router{}
                fmt.Printf("%s\n", r)
        }

}
