package chord

import (
	"log"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordAdmin/chordAdminModel"
	"google.golang.org/grpc"
)

var NetworkAdminId = ":9090"

type NetworkAdminConn struct {
	ClientConn  *grpc.ClientConn
	AdminClient chordAdminModel.ChordAdminClient
}

func NewconnToNetworkAdmin() *NetworkAdminConn {
	clientConn, err := grpc.Dial(NetworkAdminId, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Unable to get a connection")
	}
	adminClient := chordAdminModel.NewChordAdminClient(clientConn)

	return &NetworkAdminConn{
		ClientConn:  clientConn,
		AdminClient: adminClient,
	}
}

func (networkAdminConn *NetworkAdminConn) CloseConn() {
	networkAdminConn.ClientConn.Close()
}
