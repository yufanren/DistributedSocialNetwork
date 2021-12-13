package etcdclient

import "time"

const (
	dialTimeout = 5 * time.Second
	requestTimeout = 10 * time.Second
	// port of etcd server
	port = ":8013"
	// port of auth server
	authPort = ":8010"
)

func endpoints() []string {
	return []string{"localhost:2379", "localhost:22379", "localhost:32379"}
}
