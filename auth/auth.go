package auth

import (
	proto "cs9223-final-project/auth/authpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Start() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAuthServer(grpcServer, &authService{})
	log.Printf("Server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}