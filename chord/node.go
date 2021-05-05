package chord

import (
	"context"
	"log"
	"net"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/nodeModel"
)

func NewNode(id string) *Node {
	listenr, err := net.Listen("tcp", id)
	if err != nil {
		log.Fatal("Unable to listen ", err)
	}
	return &Node{
		Id:      id,
		Listner: listenr,
	}
}

func (node *Node) Ping(ctx context.Context, req *nodeModel.Request) (*nodeModel.Response, error) {
	return &nodeModel.Response{Message: "Hello World!"}, nil
}
