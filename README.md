## 1. Folder Description
* `auth`: auth server, provide token related RPC functions, run on port 8010.
* `backend`: backend server, provide user and blog related RPC functions, run on port 8011.
* `web`: web server, provide REST APIs for frontend while communicate with auth and backend servers via RPC, run on port 8012.
* `etcd_client`: etcd client, currently test only. Used in stage 2 to replace `backend` server(or only the storage part(`backend/db.go`)). Should also be a RPC server in the future.
* `cmd`: project entry point, used to start servers(users and blogs data can also be initialized here in the future).
* `Makefile`: automate build process for servers and protocol buffers. 
* `frontend`: frontend server, where directly interact with users, run on port 8080.

## 2. File Description
* `config.go`: configuration file for server arguments.
* `dial.go`: create client connections to all required RPC servers. 
* `service.go`: implement RPC functions.
* `controller.go`: implement handling functions for REST requests.
* `{same as folder name}.go`: implement the start of server.

## 3. Run Project
```shell
# terminal 0
make run-etcd
# terminal 1
make run-auth
# terminal 2
make run-etcdclient
# terminal 3
make run-web
# terminal 4
cd frontend
npm install
npm run serve
# then go to localhost:8080
```