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
		NodeDetails: &chordNodeModel.Node{
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
	// Copy data from node.NodeDetails to predecessorNodeDetails
	// This the best way I know how to do it and if anyone knows any better way please let me know
	predecessorNodeDetails := new(chordNodeModel.Node)
	predecessorNodeDetails.HashId = node.NodeDetails.HashId
	predecessorNodeDetails.Id = node.NodeDetails.Id
	node = FindNode(predecessorNodeDetails)

	for !(nodeDetails.HashId > node.NodeDetails.HashId && nodeDetails.HashId <= node.Successor.NodeDetails.HashId) {
		// Find node that has the node details of predecessorNodeDetails
		predecessorNodeDetails, err := node.FindClosestPrecedingFinger(ctx, nodeDetails)
		if err != nil {
			log.Fatal(err)
		}
		node = FindNode(predecessorNodeDetails)
		// // If i did not add the below piece of code it threw error and this is the only workaround
		// // that worked. If anyone knows a better way to handle it please suggest
		// _ = foundNode
	}

	return predecessorNodeDetails, nil
}

func FindNode(nodeDetails *chordNodeModel.Node) *Node {
	return &Node{}
}

func (node *Node) FindClosestPrecedingFinger(ctx context.Context, nodeDetails *chordNodeModel.Node) (*chordNodeModel.Node, error) {
	for i := int(HashIdSize) - 1; i >= 0; i-- {
		fingerEntryNodeDetails := node.FingerTable[i].Node.NodeDetails
		if fingerEntryNodeDetails.HashId > node.NodeDetails.HashId && fingerEntryNodeDetails.HashId <= nodeDetails.HashId {
			return node.FingerTable[i].Node.NodeDetails, nil
		}
	}
	return node.NodeDetails, nil
}
