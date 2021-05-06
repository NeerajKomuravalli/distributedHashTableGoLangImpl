package chord

import (
	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordNode/chordNodeModel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (node *Node) ServeGrpc() {
	gs := grpc.NewServer()

	chordNodeModel.RegisterChordServer(gs, node)

	reflection.Register(gs)

	node.Server = gs

	go gs.Serve(node.Listner)
}
