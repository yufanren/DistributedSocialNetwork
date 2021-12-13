package web

import (
	proto "cs9223-final-project/auth/authpb"
	"cs9223-final-project/etcdclient/etcdclientpb"
	"google.golang.org/grpc"
	"log"
)

var authClient proto.AuthClient
// var backendClient backendpb.BackendClient
var etcdClient etcdclientpb.EtcdclientClient

func dialAll() {
	dialAuth()
	// dialBackend()
	dialEtcdclient()
}

func dialAuth() {
	authConn, err := grpc.Dial(authPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect auth server: %v", err)
	}
	authClient = proto.NewAuthClient(authConn)
}

//func dialBackend() {
//	backendConn, err := grpc.Dial(backendPort, grpc.WithInsecure(), grpc.WithBlock())
//	if err != nil {
//		log.Fatalf("failed to connect backend server: %v", err)
//	}
//	backendClient = backendpb.NewBackendClient(backendConn)
//}

func dialEtcdclient() {
	etcdclientConn, err := grpc.Dial(etcdclientPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect backend server: %v", err)
	}
	etcdClient = etcdclientpb.NewEtcdclientClient(etcdclientConn)
}
