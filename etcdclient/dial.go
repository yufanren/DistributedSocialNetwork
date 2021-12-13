package etcdclient

import (
	proto "cs9223-final-project/auth/authpb"
	"google.golang.org/grpc"
	"log"
)

var authClient proto.AuthClient

func dialAll() {
	dialAuth()
}

func dialAuth() {
	authConn, err := grpc.Dial(authPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect auth server: %v", err)
	}
	authClient = proto.NewAuthClient(authConn)
}
