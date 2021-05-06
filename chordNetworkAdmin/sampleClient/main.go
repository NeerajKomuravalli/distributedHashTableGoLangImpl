package main

import (
	"context"
	"fmt"
	"log"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordAdmin/chordAdminModel"
	"google.golang.org/grpc"
)

func main() {
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
