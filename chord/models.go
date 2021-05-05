package chord

import (
	"net"

	"google.golang.org/grpc"
)

type FingerEntry struct {
	Start uint32
	Node  *Node
}

type FingerTable []FingerEntry

type Node struct {
	Successor     *Node
	FingerTable   FingerTable
	Predecessor   *Node
	Id            string
	HashId        uint32
	HashIdFloat64 float64 // This is only to avoid repitative convertion to float64 for math calculations
	Listner       net.Listener
	Server        *grpc.Server
}
