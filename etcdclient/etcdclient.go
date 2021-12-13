package etcdclient

import (
	"cs9223-final-project/etcdclient/etcdclientpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"fmt"
)

func Start() {
	dialAll()
	listener, err := net.Listen("tcp", port)
	fmt.Printf("listened")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	fmt.Printf("got server")
	etcdclientpb.RegisterEtcdclientServer(grpcServer, newEtcdService())
	go StartDB()
	log.Printf("Server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
