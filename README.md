## File Structure
```txt
📦root
 ┣ 📂auth --> auth service
 ┃ ┣ 📂authpb --> protocol buffers for auth service 
 ┃ ┃ ┣ 📜models.pb.go
 ┃ ┃ ┣ 📜models.proto
 ┃ ┃ ┗ 📜models_grpc.pb.go
 ┃ ┣ 📜auth.go --> start point for auth service
 ┃ ┣ 📜config.go --> port, token validity duration, authentication code
 ┃ ┣ 📜jwt.go --> implmentation of jwt
 ┃ ┗ 📜service.go --> implementation of auth gRPC methods
 ┣ 📂backend --> backend service(used in stage 1)
 ┃ ┣ 📂backendpb --> protocol buffers for backend service
 ┃ ┃ ┣ 📜models.pb.go
 ┃ ┃ ┣ 📜models.proto
 ┃ ┃ ┗ 📜models_grpc.pb.go
 ┃ ┣ 📜backend.go --> start point for backend service
 ┃ ┣ 📜backend_suite_test.go
 ┃ ┣ 📜config.go --> port
 ┃ ┣ 📜db.go --> storage part
 ┃ ┣ 📜db_test.go --> test for db
 ┃ ┣ 📜dial.go --> dial other gRPC servers
 ┃ ┗ 📜service.go --> implementation of backend gRPC methods
 ┣ 📂cmd --> service entry point
 ┃ ┣ 📂auth
 ┃ ┃ ┗ 📜auth.go --> start of auth service
 ┃ ┣ 📂backend
 ┃ ┃ ┗ 📜backend.go --> start of backend service
 ┃ ┣ 📂etcdclient
 ┃ ┃ ┗ 📜etcdclient.go --> start of etcdclient service
 ┃ ┗ 📂web
 ┃ ┃ ┗ 📜web.go --> start of web service
 ┣ 📂etcd --> etcd raft cluster
 ┃ ┗ 📜Procfile --> cluster set up commands
 ┣ 📂etcdclient --> backend wrapper for etcd/clientv3(used in stage 2)
 ┃ ┣ 📂etcdclientpb --> protocol buffers for etcdclient
 ┃ ┃ ┣ 📜models.pb.go
 ┃ ┃ ┣ 📜models.proto
 ┃ ┃ ┗ 📜models_grpc.pb.go
 ┃ ┣ 📜blog.go --> blog storage
 ┃ ┣ 📜blog_test.go --> test for blog
 ┃ ┣ 📜config.go --> ports and timeout
 ┃ ┣ 📜dial.go --> dial other gRPC services
 ┃ ┣ 📜etcdclient.go --> start point for etcdclient
 ┃ ┣ 📜etcdclient_suite_test.go
 ┃ ┣ 📜service.go --> implementation of etcdclient gRPC methods
 ┃ ┣ 📜user.go --> user storage
 ┃ ┣ 📜user_test.go --> test for user
 ┃ ┗ 📜utils.go --> common method utils
 ┣ 📂frontend --> frontend server
 ┃ ┣ 📂src
 ┃ ┃ ┣ 📂assets
 ┃ ┃ ┃ ┣ 📜index.css
 ┃ ┃ ┃ ┗ 📜logo.png
 ┃ ┃ ┣ 📂components
 ┃ ┃ ┃ ┗ 📜Header.vue --> page header
 ┃ ┃ ┣ 📂plugins
 ┃ ┃ ┃ ┣ 📜axios.js --> interceptors for axios
 ┃ ┃ ┃ ┗ 📜element.js --> configuration for element-ui
 ┃ ┃ ┣ 📂router
 ┃ ┃ ┃ ┗ 📜index.js --> APIs
 ┃ ┃ ┣ 📂store
 ┃ ┃ ┃ ┗ 📜index.js --> token storage
 ┃ ┃ ┣ 📂views
 ┃ ┃ ┃ ┣ 📂blog
 ┃ ┃ ┃ ┃ ┣ 📜Detail.vue --> blog detail information
 ┃ ┃ ┃ ┃ ┣ 📜Edit.vue --> blog edit panel
 ┃ ┃ ┃ ┃ ┗ 📜List.vue --> conditional blog list
 ┃ ┃ ┃ ┣ 📂sidebar
 ┃ ┃ ┃ ┃ ┣ 📜Index.vue --> index for sidebar
 ┃ ┃ ┃ ┃ ┣ 📜Promotion.vue --> auxiliary part
 ┃ ┃ ┃ ┃ ┗ 📜Welcome.vue --> post blog / login, register panel
 ┃ ┃ ┃ ┣ 📂user
 ┃ ┃ ┃ ┃ ┣ 📜Login.vue --> user login
 ┃ ┃ ┃ ┃ ┣ 📜Profile.vue --> user profile
 ┃ ┃ ┃ ┃ ┗ 📜Register.vue --> user register
 ┃ ┃ ┃ ┗ 📜Home.vue --> main page
 ┃ ┃ ┣ 📜App.vue --> root of frontend
 ┃ ┃ ┗ 📜main.js --> initialization of elements
 ┃ ┣ 📜package-lock.json
 ┃ ┗ 📜package.json --> metadata of frontend
 ┣ 📂utils --> common method utils
 ┃ ┣ 📜encoding.go --> encode/decode for object
 ┃ ┣ 📜encoding_test.go --> test for encoding
 ┃ ┗ 📜utils_suite_test.go
 ┣ 📂web --> stateless web server
 ┃ ┣ 📜config.go --> ports
 ┃ ┣ 📜controller.go --> handling functions for REST requests
 ┃ ┣ 📜dial.go --> dial other gRPC servers
 ┃ ┣ 📜result.go --> REST response template
 ┃ ┗ 📜web.go --> start point for web
 ┣ 📜.gitignore
 ┣ 📜Makefile --> service start and protocol buffers generation
 ┣ 📜README.md
 ┣ 📜gen-proto-go.sh
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┗ 📜runall --> commands to start/stop all services
```
