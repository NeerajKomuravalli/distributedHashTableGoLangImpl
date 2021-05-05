package chord

import (
	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/nodeModel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (node *Node) ServeGrpc() {
	gs := grpc.NewServer()

	nodeModel.RegisterChordServer(gs, node)

	reflection.Register(gs)

	node.Server = gs

	go gs.Serve(node.Listner)
}
