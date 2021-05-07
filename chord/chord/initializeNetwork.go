package chord

import (
	"context"
	"log"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordAdmin/chordAdminModel"
)

func InitializeNetwork(id string) *Node {
	networkConn := NewconnToNetworkAdmin()
	defer networkConn.CloseConn()

	// When we are initializing the network we have to be sure that there is no node present
	totalNoOfActiveNodes, err := networkConn.AdminClient.TotalNumberOfActiveNodes(context.Background(), &chordAdminModel.NumberOfActiveNodes{})
	if err != nil {
		log.Fatal(err)
	}
	if totalNoOfActiveNodes.Number > 0 {
		activeNodes, err := networkConn.AdminClient.GetActiveNodes(context.Background(), &chordAdminModel.ActiveNodesRequest{})
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal("Can't initialize the network as the network already has nodes present : ", activeNodes.Nodes)
	}
	adminNode := chordAdminModel.Node{Id: id}
	success, err := networkConn.AdminClient.AddActiveNode(context.Background(), &adminNode)
	if err != nil {
		log.Fatal(err)
	}
	if !success.Success {
		log.Fatal("Not able to add node to the network")
	}
	node := NewNode(id)
	node.Successor = node
	node.Predecessor = node

	return node
}
