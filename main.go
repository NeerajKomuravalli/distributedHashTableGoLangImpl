package main

import (
	"context"
	"fmt"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/chord"
	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordModel"
)

func main() {
	node1 := chord.InitializeNetwork(":8080")
	defer node1.Server.Stop()
	node2 := chord.InitializeNetwork(":8081")
	defer node2.Server.Stop()

	conn := node2.DialGrpc(node1.Id)
	defer conn.Close()
	client := chordModel.NewChordClient(conn)
	resp, err := client.Ping(context.Background(), &chordModel.Request{})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
}
