package graph

//import (
//	"github.com/justinbarrick/hone/pkg/utils"
//)

type Node interface {
	GetName() (string)
	ID() (int64)
}

//func ID(node Node) int64 {
//	return utils.Crc(node.GetName())
//}
