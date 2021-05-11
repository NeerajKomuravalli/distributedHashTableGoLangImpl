package chord

import (
	"net"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordNode/chordNodeModel"
	"google.golang.org/grpc"
)

type FingerEntry struct {
	Start uint32
	Node  *Node
}

type FingerTable []FingerEntry

type Node struct {
	NodeDetails   *chordNodeModel.Node
	Successor     *Node
	FingerTable   FingerTable
	Predecessor   *Node
	HashIdFloat64 float64 // This is only to avoid repitative convertion to float64 for math calculations
	Listner       net.Listener
	Server        *grpc.Server
}
