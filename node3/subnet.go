package main

type Subnet struct {
  Name             string           `yaml:"name"`
  Netaddr           string          `yaml:"netaddr"`
  _ID   int64
}

//type Subnet struct {
//	_Name string
//	_ID   int64
//}

func (j Subnet) GetName() string {
	//return j._Name
	return j.Name
}
func (j Subnet) ID() int64 {
	return j._ID
}
