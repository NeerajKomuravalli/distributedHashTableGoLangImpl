package chordNetworkAdminServer

import (
	"context"
	"fmt"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordAdmin/chordAdminModel"
	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordNode/chordNodeModel"
)

type NetworkAdminServer struct{}

func NewNetworkAdminServer() *NetworkAdminServer {
	return &NetworkAdminServer{}
}

var ActiveNodes = new(chordAdminModel.ActiveNodes)

func (networkServer *NetworkAdminServer) AddActiveNode(ctx context.Context, node *chordNodeModel.Node) (*chordAdminModel.Success, error) {
	// Check if the mode exists and if it does then return success as false
	found, _ := search(ActiveNodes, node.Id)
	if found {
		return &chordAdminModel.Success{Success: false}, fmt.Errorf("node already present")
	}
	ActiveNodes.Nodes = append(ActiveNodes.Nodes, node)
	return &chordAdminModel.Success{Success: true}, nil
}

// Search in ActiveNodes.Nodes to make sure same ip is not added twice
// And yes, this is a linear search and also known as the worst search possible.
// I am doing this to save time and will readdress this to make it optimal when
// I have time
func search(activeNodes *chordAdminModel.ActiveNodes, id string) (bool, int) {
	for index, v := range activeNodes.Nodes {
		if v.Id == id {
			return true, index
		}
	}
	return false, len(activeNodes.Nodes)
}

func (networkServer *NetworkAdminServer) GetActiveNodes(ctx context.Context, req *chordAdminModel.ActiveNodesRequest) (*chordAdminModel.ActiveNodes, error) {
	return ActiveNodes, nil
}

func (networkServer *NetworkAdminServer) TotalNumberOfActiveNodes(ctx context.Context, req *chordAdminModel.NumberOfActiveNodes) (*chordAdminModel.TotalNumberOfNodes, error) {
	TotalNumberOfNodes := new(chordAdminModel.TotalNumberOfNodes)
	TotalNumberOfNodes.Number = uint64(len(ActiveNodes.Nodes))
	return TotalNumberOfNodes, nil
}
