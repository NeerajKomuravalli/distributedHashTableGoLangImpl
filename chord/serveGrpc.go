package chord

import (
	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordModel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (node *Node) ServeGrpc() {
	gs := grpc.NewServer()

	chordModel.RegisterChordServer(gs, node)

	reflection.Register(gs)

	node.Server = gs

	go gs.Serve(node.Listner)
}
