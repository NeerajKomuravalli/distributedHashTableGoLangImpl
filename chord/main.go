package main

import (
	"context"
	"fmt"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/chord/chord"
	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordNode/chordNodeModel"
)

func main() {
	node1 := chord.NewNode(":8080")
	defer node1.Server.Stop()

	node2 := chord.NewNode(":8081")
	defer node2.Server.Stop()

	conn := node2.DialGrpc(node1.Id)
	defer conn.Close()
	client := chordNodeModel.NewChordClient(conn)
	resp, err := client.Ping(context.Background(), &chordNodeModel.Request{})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
}
