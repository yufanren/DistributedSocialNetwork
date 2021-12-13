package backend

import (
	"final-project-0b5a2e16-ly2062-yufanren/backend/backendpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Start() {
	dialAll()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	backendpb.RegisterBackendServer(grpcServer, newBackendService())
	log.Printf("Server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
