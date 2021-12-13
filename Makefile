AUTH_PB_PATH=auth/authpb/models.proto
BACKEND_PB_PATH=backend/backendpb/models.proto
ETCDCLIENT_PB_PATH=etcdclient/etcdclientpb/models.proto
WEB_ENTRY_POINT=cmd/web/web.go
AUTH_ENTRY_POINT=cmd/auth/auth.go
BACKEND_ENTRY_POINT=cmd/backend/backend.go
ETCDCLIENT_ENTRY_POINT=cmd/etcdclient/etcdclient.go

protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(AUTH_PB_PATH)
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(BACKEND_PB_PATH)
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(ETCDCLIENT_PB_PATH)

run-web:
	go run $(WEB_ENTRY_POINT)

run-auth:
	go run $(AUTH_ENTRY_POINT)

run-backend:
	go run $(BACKEND_ENTRY_POINT)

run-etcdclient:
	go run $(ETCDCLIENT_ENTRY_POINT)

run-etcd:
	etcd

# to do
# 1. how to run the commands in background so that they can run together without blocking?
# 2. how to stop server elegantly?