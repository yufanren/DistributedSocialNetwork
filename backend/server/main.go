package main

import (
	"context"
	"dist-demo/auth"
	"dist-demo/proto"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8090"
)

type server struct {
	proto.UnimplementedDemoServer
	db *DB
}

func newServer() *server {
	s := server{}
	s.db = newDB()
	return &s
}

func (s *server) Login(ctx context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {
	if !s.db.hasUser(req.Username) {
		return &proto.AuthResponse{Token: ""}, errors.New("user does not exist")
	}
	if !s.db.validUser(req.Username, req.Password) {
		return &proto.AuthResponse{Token: ""}, errors.New("wrong password")
	}
	token, _ := auth.GetToken(req.Username)
	return &proto.AuthResponse{Token: token}, nil
}

func (s *server) Register(ctx context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {
	if err := s.db.addUser(req.Username, req.Password); err != nil {
		return &proto.AuthResponse{Token: ""}, err
	}
	token, _ := auth.GetToken(req.Username)
	return &proto.AuthResponse{Token: token}, nil
}

func (s *server) ListBlog(ctx context.Context, req *proto.ListBlogRequest) (*proto.ListBlogResponse, error) {
	total, blogs := s.db.listBlogs(int(req.PageNum), int(req.PageSize))
	return &proto.ListBlogResponse{Total: int32(total), Blogs: blogs}, nil
}

func (s *server) AddBlog(ctx context.Context, req *proto.AddBlogRequest) (*proto.AddBlogResponse, error) {
	return &proto.AddBlogResponse{BlogId: int32(s.db.addBlog(req.Blog))}, nil
}

func (s *server) GetBlog(ctx context.Context, req *proto.GetBlogRequest) (*proto.GetBlogResponse, error) {
	return &proto.GetBlogResponse{Blog: s.db.getBlog(int(req.BlogId))}, nil
}

func (s *server) Follow(ctx context.Context, req *proto.FollowRequest) (*proto.Response, error) {
	s.db.followUser(req.Parent, req.Follower)
	return &proto.Response{}, nil
}

func (s *server) IsFollow(ctx context.Context, req *proto.FollowRequest) (*proto.CheckResponse, error) {
	return &proto.CheckResponse{Res: s.db.isFollow(req.Parent, req.Follower)}, nil
}

func (s *server) GetUserBlog(ctx context.Context, req *proto.UserBlogRequest) (*proto.ListBlogResponse, error) {
	total, blogs := s.db.listBlogsFromUser(string(req.Username), int(req.PageNum), int(req.PageSize))
	return &proto.ListBlogResponse{Total: int32(total), Blogs: blogs}, nil
}

func (s *server) GetUserFollowed(ctx context.Context, req *proto.UserRequest) (*proto.ListUserResponse, error) {
	total, users := s.db.listFollowedFromUser(string(req.Username))
	return &proto.ListUserResponse{Total: int32(total), Parents: users}, nil
}

func (s *server) GetUserFollower(ctx context.Context, req *proto.UserRequest) (*proto.ListUserResponse, error) {
	total, users := s.db.listFollowerFromUser(string(req.Username))
	return &proto.ListUserResponse{Total: int32(total), Parents: users}, nil
}

func (s *server) GetFollowBlog(ctx context.Context, req *proto.UserBlogRequest) (*proto.ListBlogResponse, error) {
	total, blogs := s.db.listFollowBlog(string(req.Username), int(req.PageNum), int(req.PageSize))
	return &proto.ListBlogResponse{Total: int32(total), Blogs: blogs}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	s := newServer()
	proto.RegisterDemoServer(grpcServer, s)
	log.Printf("Server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}