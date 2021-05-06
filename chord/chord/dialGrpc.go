package chord

import (
	"log"

	"google.golang.org/grpc"
)

func (node *Node) DialGrpc(id string) *grpc.ClientConn {
	conn, err := grpc.Dial(id, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Unable to get a connection ", err)
	}
	return conn
}
