package main



import (
  "local.packages/graph"
  "fmt"
  "net"
  //"strings"
  "io/ioutil"
  //yaml "gopkg.in/yaml.v2"
  yaml    "github.com/goccy/go-yaml"
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
  Routers  []Router `yaml:"routers"`
  Subnets  []Subnet `yaml:"subnets"`
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
  fmt.Printf("%v\n", router_dic)
  fmt.Printf("%v\n", subnet_dic)

  err = yaml.Unmarshal(buf, &d)
  if err != nil {
    panic(err)
  }

/*
  for _, v := range d.Routers {

    fmt.Printf("%s\n", v.HostName)
    router_dic[v.HostName] = v
    for _, i := range v.Interfaces {
        fmt.Printf("    %s\t%s\n", i.Name, i.Ipaddr)
    }

  }

  for _, v := range d.Subnets {

    fmt.Printf("%s\t%s\n", v.Name, v.Netaddr)
    //subnet_dic[v.Name] = v
    subnet_dic[v.Netaddr] = v

  }
  fmt.Printf("----------------------------------------\n")
  fmt.Printf("router_dic\n%v\n\n", router_dic)
  fmt.Printf("subnet_dic\n%v\n\n", subnet_dic)
  fmt.Printf("----------------------------------------\n")
*/
/*
  nodes := []graph.Node{}
//  router :=  Router{ "routerA",1 }
//  router2 :=  Router{ "routerB",2 }
//  subnet :=  Subnet{ "subnet1", 10}

  router :=  d.Routers[0]
  nodes = append(nodes, router)
*/
 var id int64 
 id = 1
  nodes := []graph.Node{}
  g := graph.NewGraph(nodes)
  for _, v := range d.Routers {

    //fmt.Printf("%s\n", v.HostName)
    v._ID = id
    id++
    nodes = append(nodes, v)
    g.AddNode(v)
    router_dic[v.HostName] = v

  }

  for _, v := range d.Subnets {

    v._ID = id
    id++
    nodes = append(nodes, v)
    g.AddNode(v)
    subnet_dic[v.Netaddr] = v

  }
  fmt.Printf("----------------------------------------\n")
  fmt.Printf("router_dic\n%v\n\n", router_dic)
  fmt.Printf("subnet_dic\n%v\n\n", subnet_dic)
  fmt.Printf("----------------------------------------\n")
  // := graph.NewGraph(nodes)
  //g := graph.NewGraph(nodes)

  //g.Dump()
/*
  fmt.Printf("----------------------------------------\n")
  for _, v := range d.Routers {

    fmt.Printf("%s\n", v.HostName)
    for _, i := range v.Interfaces {
        fmt.Printf("    %s\t%s\n", i.Name, i.Ipaddr)
	//ip, ipnet, _ := net.ParseCIDR(i.Ipaddr)
	_, ipnet, _ := net.ParseCIDR(i.Ipaddr)
        masklen, _ := ipnet.Mask.Size() 
        //fmt.Println(ip) // 
        //fmt.Println(ipnet.IP) // 10.0.0.0
        //fmt.Println(ipnet.Mask) // ffffff00
        //fmt.Println(masklen) // 24
        subnet := fmt.Sprintf("%s/%d", ipnet.IP, masklen)
        fmt.Println(subnet) // 
        s := subnet_dic[subnet]
        fmt.Printf("r   %v\n", v)
        fmt.Printf("s   %v\n", s)
        //g.SetEdge(v, s)

    }

  }
*/

  fmt.Printf("----------------------------------------\n")
  //for _, v := range d.Routers {
  //for k, v := range router_dic {
  for _, v := range router_dic {

    fmt.Printf("%s\n", v.HostName)
    for _, i := range v.Interfaces {
        fmt.Printf("    %s\t%s\n", i.Name, i.Ipaddr)
	//ip, ipnet, _ := net.ParseCIDR(i.Ipaddr)
	_, ipnet, _ := net.ParseCIDR(i.Ipaddr)
        masklen, _ := ipnet.Mask.Size() 
        //fmt.Println(ip) // 
        //fmt.Println(ipnet.IP) // 10.0.0.0
        //fmt.Println(ipnet.Mask) // ffffff00
        //fmt.Println(masklen) // 24
        subnet := fmt.Sprintf("%s/%d", ipnet.IP, masklen)
        fmt.Println(subnet) // 
        s := subnet_dic[subnet]
        fmt.Printf("r   %v\n", v)
        fmt.Printf("s   %v\n", s)
        g.SetEdge(v, s)

    }

  }
  //g.SetEdge(d.Routers[0], d.Subnets[0])

  g.Dump()




/*
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
// g.Dump(s)
*/




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
