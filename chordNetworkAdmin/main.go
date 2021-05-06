package main

import (
	"log"
	"net"

	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/chordNetworkAdmin/chordNetworkAdminServer"
	"github.com/NeerajKomuravalli/distributedHashTableGoLangImpl/models/proto/chordAdmin/chordAdminModel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var ipAddress = ":9090"

func main() {
	gs := grpc.NewServer()

	server := chordNetworkAdminServer.NewNetworkAdminServer()

	// register the temepfrature conversion server
	chordAdminModel.RegisterChordAdminServer(gs, server)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	listen, err := net.Listen("tcp", ipAddress)
	if err != nil {
		log.Fatal("Unable to listen")
	}
	gs.Serve(listen)
}
