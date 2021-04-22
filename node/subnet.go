package main

type Subnet struct {
	_Name string
	_ID   int64
}

func (j Subnet) GetName() string {
	return j._Name
}
func (j Subnet) ID() int64 {
	return j._ID
}
