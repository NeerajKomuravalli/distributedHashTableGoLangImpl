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
		NodeDetails: chordNodeModel.Node{
			Id:     "id",
			HashId: GenerateHashId(id),
		},
		Listner:     listenr,
		Successor:   nil,
		Predecessor: nil,
	}

	// This is done to make mathematical operations optimal. This avoids repeated conversion
	node.HashIdFloat64 = float64(node.NodeDetails.HashId)
	node.GenerateNewFingerTable()
	node.ServeGrpc()
	return &node
}

func (node *Node) Ping(ctx context.Context, req *chordNodeModel.Request) (*chordNodeModel.Response, error) {
	return &chordNodeModel.Response{Message: "Hello World!"}, nil
}

func (node *Node) FindSuccessor(ctx context.Context, nodeDetails *chordNodeModel.Node) (*chordNodeModel.Node, error) {
	return nodeDetails, nil
}

func (node *Node) FindPredecessor(ctx context.Context, nodeDetails *chordNodeModel.Node) (*chordNodeModel.Node, error) {
	return nodeDetails, nil
}

func (node *Node) FindClosestPrecedingFinger(ctx context.Context, nodeDetails *chordNodeModel.Node) (*chordNodeModel.Node, error) {
	for i := int(HashIdSize) - 1; i >= 0; i-- {
		// finger := node.FingerTable[i]
	}
	return nodeDetails, nil
}
