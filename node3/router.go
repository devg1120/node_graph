package main

type Router struct {
  HostName         string           `yaml:"hostname"`
  Interfaces       []Interface      `yaml:"interfaces"`
  _ID   int64
}

type Interface struct {
  Name         string           `yaml:"name"`
  Ipaddr       string           `yaml:"ipaddr"`
}

//type Router struct {
//	_Name string
//	_ID   int64
//}


func (j Router) GetName() string {
	//return j._Name
	return j.HostName
}

func (j Router) ID() int64 {
	return j._ID
}

