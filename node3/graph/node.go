package graph

//import (
//	"github.com/justinbarrick/hone/pkg/utils"
//)

type NetNode interface {
	GetName() string
	ID() int64
}

//func ID(node Node) int64 {
//	return utils.Crc(node.GetName())
//}
