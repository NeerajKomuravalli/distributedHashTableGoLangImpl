package chordNetworkAdminServer

import (
	"context"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordAdmin/chordAdminModel"
)

type NetworkAdminServer struct{}

func NewNetworkAdminServer() *NetworkAdminServer {
	return &NetworkAdminServer{}
}

var ActiveNodes = new(chordAdminModel.ActiveNodes)

func (networkServer *NetworkAdminServer) AddActiveNode(ctx context.Context, node *chordAdminModel.Node) (*chordAdminModel.Success, error) {
	ActiveNodes.Nodes = append(ActiveNodes.Nodes, node)
	return &chordAdminModel.Success{Success: true}, nil
}

func (networkServer *NetworkAdminServer) GetActiveNodes(ctx context.Context, req *chordAdminModel.ActiveNodesRequest) (*chordAdminModel.ActiveNodes, error) {
	return ActiveNodes, nil
}
