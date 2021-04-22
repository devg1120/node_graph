package main

import (
  "fmt"
  //"strings"
  "io/ioutil"
  //yaml "gopkg.in/yaml.v2"
  yaml    "github.com/goccy/go-yaml"
  //  "gopkg.in/go-playground/validator.v9"

)

//----------------------------------------------------
type Data struct {
  Routers  []Router `yaml:"routers"`
  Subnets  []Subnet `yaml:"subnets"`
}

type Router struct {
  HostName         string           `yaml:"hostname"`
  Interfaces       []Interface      `yaml:"interfaces"`
}

type Interface struct {
  Name         string           `yaml:"name"`
  Ipaddr       string           `yaml:"ipaddr"`
}

type Subnet struct {
  Name             string           `yaml:"name"`
  Netaddr           string          `yaml:"netaddr"`
}

//----------------------------------------------------
/*
type Data struct {
  Users  []User `yaml:"users"`
  Entrys []User `yaml:"entrys"`
}

type User struct {
  Name             string           `yaml:"common"`
  FullName         fullName         `yaml:"full_name"`
  Sex              string           `yaml:"sex"`
  SelfIntroduction selfIntroduction `yaml:"self_introduction"`
  ImageURLs        []string         `yaml:"image_urls"`
  Shemale          bool             `yaml:"shemale"`
}

type fullName struct {
  FirstName string `yaml:"first_name"`
  LastName  string `yaml:"last_name"`
}

type selfIntroduction struct {
  Long  string `yaml:"long"`
  Short string `yaml:"short"`
}
*/

func main() {

  buf, err := ioutil.ReadFile("./network.yaml")
  if err != nil {
    panic(err)
  }
  //fmt.Printf("buf: %+v\n", string(buf))

  var d Data
  err = yaml.Unmarshal(buf, &d)
  if err != nil {
    panic(err)
  }

  //fmt.Printf("d: %+v\n\n", d)
  //fmt.Printf("d: %+v\n\n", d.Routers[0])
  //fmt.Printf("d: %+v\n\n", d.Subnets[1])

  for _, v := range d.Routers {

    fmt.Printf("%s\n", v.HostName)
    for _, i := range v.Interfaces {
        fmt.Printf("    %s\t%s\n", i.Name, i.Ipaddr)
    }

  }

}
