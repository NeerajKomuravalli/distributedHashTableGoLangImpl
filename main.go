package main

import (
	"context"
	"fmt"
	"log"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/chord"
	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/chordNetworkAdmin"
	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordAdmin/chordAdminModel"
	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordNode/chordNodeModel"
	"google.golang.org/grpc"
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

	grpcServer := chordNetworkAdmin.ServeChordNetworkAdmin(":9090")
	defer grpcServer.Stop()
	clientConn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Unable to get a connection")
	}
	defer clientConn.Close()
	adminClient := chordAdminModel.NewChordAdminClient(clientConn)
	activeNode, err := adminClient.GetActiveNodes(context.Background(), new(chordAdminModel.ActiveNodesRequest))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Before adding any node : %v\n", activeNode)
	node := chordAdminModel.Node{Id: "22"}
	success, err := adminClient.AddActiveNode(context.Background(), &node)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Adding of node was success : %v\n", success.Success)
	node = chordAdminModel.Node{Id: "102"}
	success, err = adminClient.AddActiveNode(context.Background(), &node)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Adding of node was success : %v\n", success.Success)
	activeNodeLater, err := adminClient.GetActiveNodes(context.Background(), new(chordAdminModel.ActiveNodesRequest))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After adding node : %v\n", activeNodeLater)
}
