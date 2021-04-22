package main

type Router struct {
	_Name string
	_ID   int64
}

func (j Router) GetName() string {
	return j._Name
}
func (j Router) ID() int64 {
	return j._ID
}
