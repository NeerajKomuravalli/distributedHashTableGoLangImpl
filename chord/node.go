package chord

import (
	"context"
	"log"
	"net"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordNode/chordNodeModel"
)

func NewNode(id string) *Node {
	listenr, err := net.Listen("tcp", id)
	if err != nil {
		log.Fatal("Unable to listen ", err)
	}
	node := Node{
		Id:          id,
		Listner:     listenr,
		HashId:      GenerateHashId(id),
		Successor:   nil,
		Predecessor: nil,
	}

	node.HashIdFloat64 = float64(node.HashId)
	node.GenerateNewFingerTable()
	node.ServeGrpc()
	return &node
}

func (node *Node) Ping(ctx context.Context, req *chordNodeModel.Request) (*chordNodeModel.Response, error) {
	return &chordNodeModel.Response{Message: "Hello World!"}, nil
}

func (node *Node) FindSuccessor(id string) *Node {
	if node.Id == id {
		return node
	} else {
		return nil
	}
}

func (node *Node) FindPredecessor(id string) *Node {
	if node.Id == id {
		return node
	} else {
		return nil
	}
}
