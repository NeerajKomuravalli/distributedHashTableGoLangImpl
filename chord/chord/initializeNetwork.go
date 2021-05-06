package chord

import (
	"context"
	"log"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordAdmin/chordAdminModel"
)

func InitializeNetwork(id string) *Node {
	networkConn := NewconnToNetworkAdmin()
	defer networkConn.CloseConn()

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
